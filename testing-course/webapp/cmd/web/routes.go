package main

import (
	"github.com/go-chi/chi/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	//Register Middleware
	mux.Use(middleware.Recoverer)

	//register routes

	//static assets

	return mux
}
