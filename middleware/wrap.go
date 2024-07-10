package middleware

import (
	"net/http"
)

func Wrap(handler http.Handler, middlewareList ...Middleware) http.Handler {
	firstMiddleware := middlewareList[0]
	restMiddleware := middlewareList[1:]

	wrappedHandler := firstMiddleware(handler)

	if len(restMiddleware) == 0 {
		return wrappedHandler
	}

	return Wrap(wrappedHandler, restMiddleware...)
}
