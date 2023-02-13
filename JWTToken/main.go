package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("super-secret-auth-key")
var apiKey = "1234"

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Minute).Unix()

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func GetJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] == apiKey {
			token, err := CreateJWT()

			if err != nil {
				return
			}

			fmt.Fprint(w, token)
		}
	}
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)

				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}

				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Area")
}

func main() {
	http.Handle("/api", ValidateJWT(Home))
	http.HandleFunc("/jwt", GetJWT)

	http.ListenAndServe(":3500", nil)
}
