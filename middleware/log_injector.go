package middleware

import (
	"net/http"

	"github.com/sparkymat/pover/log"
	"go.uber.org/zap"
)

func LogInjector(l *zap.SugaredLogger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(log.InjectIntoContext(r.Context(), l))
			next.ServeHTTP(w, r)
		})
	}
}
