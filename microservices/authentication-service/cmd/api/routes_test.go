package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"testing"
)

func test_routes_exist(t *testing.T) {
	testApp := Config{}

	testroutes := testApp.routes()
	chiRoutes := testroutes.(chi.Router)

	routes := []string{"/authenticate"}

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}

}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes,func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}

		return nil
	})

	if !found {
		t.Errorf("did not finf %s in registered routes", route)
	}
}
