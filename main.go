package main

//go:generate go run github.com/a-h/templ/cmd/templ@latest generate

import (
	"embed"
	"io/fs"
	"net/http"
	"time"

	"github.com/sparkymat/pover/internal/config"
	"github.com/sparkymat/pover/internal/handler"
	"github.com/sparkymat/pover/povc"
)

//go:embed public/css
var publicCSSFolder embed.FS

//go:embed public/js
var publicJSFolder embed.FS

//go:embed public/fonts
var publicFontsFolder embed.FS

//go:embed app/pover.rb
var poverCode []byte

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	p := povc.New(cfg, poverCode)

	cssFolder, err := fs.Sub(publicCSSFolder, "public/css")
	if err != nil {
		panic(err)
	}

	jsFolder, err := fs.Sub(publicJSFolder, "public/js")
	if err != nil {
		panic(err)
	}

	fontsFolder, err := fs.Sub(publicFontsFolder, "public/fonts")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.FS(cssFolder))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.FS(jsFolder))))
	mux.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.FS(fontsFolder))))
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(cfg.StorageFolder()))))

	mux.Handle("GET /{$}", handler.Home())
	mux.Handle("POST /generate_image", handler.GenerateImage(p))

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second, //nolint:mnd
	}

	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}
