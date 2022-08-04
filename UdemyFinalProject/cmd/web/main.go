package main

import (
	"database/sql"
	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/redisstore"
	_ "github.com/alexedwards/scs/redisstore"
	_ "github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
	"os"
	"sync"
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
		Session:  session,
		DB:       db,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Wait:     &wg,
	}
	// setup mail

	// listen for web connections
}

func initSession() *scs.SessionManager {
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
