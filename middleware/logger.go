package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		log.Println("before")
		next.ServeHTTP(w, r)
		log.Println("after")
	}

	return http.HandlerFunc(f)
}
