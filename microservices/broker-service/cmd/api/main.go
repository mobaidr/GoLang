package main

import (
	"fmt"
	"log"
	"net/http"
)

const WebPort = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting broker service on post %s", WebPort)

	//Define HTtp Server.

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.routes(),
	}

	// start the server.
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
