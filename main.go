package main

//go:generate go run github.com/a-h/templ/cmd/templ@latest generate

import (
	"net/http"
	"time"

	"github.com/sparkymat/pover/internal/config"
	"github.com/sparkymat/pover/internal/handler"
	"github.com/sparkymat/pover/povc"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	p := povc.New(cfg)

	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("public/js"))))
	mux.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("public/fonts"))))
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(cfg.StorageFolder()))))

	mux.Handle("GET /{$}", handler.Home())
	mux.Handle("POST /generate_image", handler.GenerateImage(p))

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}
