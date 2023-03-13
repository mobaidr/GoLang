package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/context"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"x-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	var app application

	//Create a dummy handler that will use to check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure that the value exists in the context
		val := r.Context().Value(contextUserKey)

		if val == nil {
			t.Error(contextUserKey, " not present")
		}

		// make sure we got a string back.
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}

		t.Log(ip)
	})

	for _, e := range tests {
		// create handler to tests
		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAddr {
			req.RemoteAddr = ""

		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFomrContext(t *testing.T) {
	var app application

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextUserKey, "whatever")

	ip := app.ipFromContext(ctx)

	if !strings.EqualFold("whatever", ip) {
		t.Error("Wrong Value returned from Context")
	}
}
