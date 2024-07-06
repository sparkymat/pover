package handler

import (
	"fmt"
	"net/http"

	"github.com/sparkymat/pover/internal/view"
)

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		component := view.Home()

		document := view.Layout("pover", component)

		err := document.Render(r.Context(), w)
		if err != nil {
			fmt.Printf("err: %v", err) //nolint:forbidigo
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
}
