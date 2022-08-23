package main

import (
	"database/sql"
	"encoding/gob"
	"finalproject/data"
	"fmt"
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const webPort = "80"

func main() {
	// Connect to a database
	db := initDB()

	// Create sessions
	session := initSession()

	//Create Logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create some channels

	// create Wait group
	wg := sync.WaitGroup{}

	// set up application config
	app := Config{
		Session:       session,
		DB:            db,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		Wait:          &wg,
		Models:        data.New(db),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	// setup mail
	app.Mailer = app.createMail()
	go app.listenForMail()

	// listen to signals
	go app.listenForShutDown()

	//listen for errors
	go app.ListenForErrors()

	// listen for web connections
	app.serve()
}

func (app *Config) ListenForErrors() {
	for {
		select {
		case err := <-app.ErrorChan:
			{
				app.ErrorLog.Println(err)
			}
		case <-app.ErrorChanDone:
			{
				return
			}
		}
	}
}

func (app *Config) serve() {
	// start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	app.InfoLog.Println("Starting Server...")

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func initSession() *scs.SessionManager {
	gob.Register(data.User{})
	//set up session
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}

	return redisPool
}

func initDB() *sql.DB {
	conn := connectToDatabase()

	if conn == nil {
		log.Panic("Can't connect to Database")
		return nil
	}

	return conn
}

func connectToDatabase() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println("Postgres not yet ready...")
		} else {
			log.Println("connected to database")

			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Println("Backing off for a second")
		time.Sleep(time.Second)
		counts++
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) listenForShutDown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.shutdown()

	os.Exit(0)
}

func (app *Config) shutdown() {
	//Perform any clean up.
	app.InfoLog.Println("would run cleanup task")

	app.Wait.Wait()

	app.Mailer.MailerDoneChan <- true
	app.ErrorChanDone <- true

	app.InfoLog.Println("Closing Channels & Shutdown application ...")

	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.MailerDoneChan)
	close(app.ErrorChan)
	close(app.ErrorChanDone)
}

func (app *Config) createMail() Mail {
	//Create Channels
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	doneChan := make(chan bool)

	m := Mail{
		Domain:         "localhost",
		Host:           "localhost",
		Port:           1025,
		Encryption:     "none",
		FromAddress:    "info@mycompany.com",
		FromName:       "Info",
		Wait:           app.Wait,
		MailerChan:     mailerChan,
		ErrorChan:      errorChan,
		MailerDoneChan: doneChan,
	}

	return m
}
