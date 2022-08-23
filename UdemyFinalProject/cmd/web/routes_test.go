package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

var routes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/members/plans",
	"/members/subscribe",
}

func Test_Routes_Exists(t *testing.T) {
	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	for _, route := range routes {
		routeExists(
			t,
			chiRoutes,
			route,
		)
	}
}

func routeExists(t *testing.T, chiRoutes chi.Router, route string) {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}

		return nil
	})

	if !found {
		t.Errorf("Did not find %s in registered routes", route)
	}
}
