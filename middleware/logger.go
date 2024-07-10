package middleware

import (
	"net/http"
	"time"

	logpkg "github.com/sparkymat/pover/log"
)

func Logger(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		log := logpkg.FromContext(r.Context())

		t := time.Now()
		next.ServeHTTP(w, r)
		log.Infof("Got %s on %s from %s. Responded in %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(t))
	}

	return http.HandlerFunc(f)
}
