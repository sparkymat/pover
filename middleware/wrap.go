package middleware

import (
	"net/http"
)

func Wrap(handler http.Handler, middlewareList ...Middleware) http.Handler {
	lastMiddleware := middlewareList[len(middlewareList)-1]
	restMiddleware := middlewareList[:len(middlewareList)-1]

	wrappedHandler := lastMiddleware(handler)

	if len(restMiddleware) == 0 {
		return wrappedHandler
	}

	return Wrap(wrappedHandler, restMiddleware...)
}
