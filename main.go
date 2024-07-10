package main

//go:generate go run github.com/a-h/templ/cmd/templ@latest generate

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/sparkymat/pover/internal/config"
	"github.com/sparkymat/pover/internal/handler"
	logpkg "github.com/sparkymat/pover/log"
	"github.com/sparkymat/pover/middleware"
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
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	logger, log, err := logpkg.Init()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	log.Info("Starting pover")

	cfg, err := config.New()
	if err != nil {
		log.Errorf("error loading configuration: %v", err)

		exitCode = 1

		return
	}

	if err = os.MkdirAll(cfg.StorageFolder(), 0o750); err != nil { //nolint:mnd
		log.Errorf("error creating storage folder: %v", err)

		exitCode = 1

		return
	}

	p := povc.New(cfg, poverCode)

	cssFolder, err := fs.Sub(publicCSSFolder, "public/css")
	if err != nil {
		log.Errorf("error loading css folder: %v", err)

		exitCode = 1

		return
	}

	jsFolder, err := fs.Sub(publicJSFolder, "public/js")
	if err != nil {
		log.Errorf("error loading js folder: %v", err)

		exitCode = 1

		return
	}

	fontsFolder, err := fs.Sub(publicFontsFolder, "public/fonts")
	if err != nil {
		log.Errorf("error loading fonts folder: %v", err)

		exitCode = 1

		return
	}

	middlewares := []middleware.Middleware{
		middleware.LogInjector(log), // note: This has to be the first one
		middleware.Logger,
	}

	wrap := func(h http.Handler) http.Handler {
		return middleware.Wrap(h, middlewares...)
	}

	mux := http.NewServeMux()
	mux.Handle("/css/", wrap(http.StripPrefix("/css/", http.FileServer(http.FS(cssFolder)))))
	mux.Handle("/js/", wrap(http.StripPrefix("/js/", http.FileServer(http.FS(jsFolder)))))
	mux.Handle("/fonts/", wrap(http.StripPrefix("/fonts/", http.FileServer(http.FS(fontsFolder)))))
	mux.Handle("/images/", wrap(http.StripPrefix("/images/", http.FileServer(http.Dir(cfg.StorageFolder())))))

	mux.Handle("GET /{$}", wrap(handler.Home()))
	mux.Handle("POST /generate_image", wrap(handler.GenerateImage(p)))

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second, //nolint:mnd
	}

	if err = server.ListenAndServe(); err != nil {
		log.Errorf("error starting server: %v", err)

		exitCode = 1

		return
	}
}
