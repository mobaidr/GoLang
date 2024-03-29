package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"
var counts int64

type Config struct {
	Repo data.Repository
	Client *http.Client
}

func main() {
	log.Println("Starting Authentication Service.")

	// Connect to Postgres Database
	conn := connectToDB()

	if conn == nil {
		log.Panic("Cannot connect to Postgres")
	}

	// Set app config
	app := Config{
		Client: &http.Client{},
	}

	app.setupRepo(conn)

	srv := http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	log.Println(dsn)
	time.Sleep(5*time.Second)

	for {
		connection ,err := openDB(dsn)

		if err != nil {
			log.Println("Postgres not yet ready..")
			log.Println(err)
			counts++
		} else {
			log.Println("Connected to Postgres")
			return  connection
		}

		if counts > 10 {
			log.Println("err")

			return nil
		}

		log.Println("Backing off for 2 seconds....")
		time.Sleep(2 * time.Second)

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

func (app *Config) setupRepo(conn *sql.DB)  {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}
