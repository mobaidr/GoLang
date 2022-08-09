package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) Routes() http.Handler {
	//Create router
	mux := chi.NewRouter()

	//Setup Middleware
	mux.Use(middleware.Recoverer)

	//Define Application routes
	mux.Get("/", app.HomePage)

	return mux
}
