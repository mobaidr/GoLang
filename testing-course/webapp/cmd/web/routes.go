package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	//Register Middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIPToContext)

	//register routes
	mux.Get("/", app.Home)

	//static assets
	filserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", filserver))

	return mux
}
