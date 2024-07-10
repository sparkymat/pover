package handler

import (
	"net/http"

	"github.com/sparkymat/pover/internal/view"
	logpkg "github.com/sparkymat/pover/log"
)

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := logpkg.FromContext(r.Context())

		component := view.Home()
		document := view.Layout("pover", component)

		err := document.Render(r.Context(), w)
		if err != nil {
			log.Errorf("error rendering document: %v", err)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
}
