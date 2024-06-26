package main

import (
	"net/http"

	"github.com/sparkymat/pover/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /", handler.Home())

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
