package middleware

import (
	"net/http"
	"time"

	logpkg "github.com/sparkymat/pover/log"
)

func Logger(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		log := logpkg.FromContext(r.Context())

		beforeTime := time.Now()
		log.Debugf("Started %s %s at %v", r.Method, r.URL.Path, beforeTime.Format(time.RFC3339Nano))

		next.ServeHTTP(w, r)

		afterTime := time.Now()
		log.Debugf("Completed %s %s at %v", r.Method, r.URL.Path, afterTime.Format(time.RFC3339Nano))

		deltaTime := time.Since(beforeTime)
		log.Infof("Got %s on %s from %s. Responded in %v", r.Method, r.URL.Path, r.RemoteAddr, deltaTime)
	}

	return http.HandlerFunc(f)
}
