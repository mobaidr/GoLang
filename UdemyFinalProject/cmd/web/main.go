package main

import (
	"database/sql"
	"log"
	"os"
	"time"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"
func main() {
	// Connect to a database
	db := initDB()
	db.Ping()

	// Create sessions

	// Create some channels

	// create Wait group

	// set up application config

	// setup mail

	// listen for web connections
}

func initDB() *sql.DB {
	conn := connectToDatabase()

	if conn == nil {
		log.Panic("Can't connect to Database")
		return nil
	}

	return conn
}

func connectToDatabase() *sql.DB  {
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
