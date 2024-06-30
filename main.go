package main

import (
	"net/http"

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

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
