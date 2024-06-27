package main

import (
	"net/http"

	"github.com/sparkymat/pover/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("public/js"))))
	mux.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("public/fonts"))))
	mux.Handle("GET /{$}", handler.Home())

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
