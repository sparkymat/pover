package handler

import (
	"context"
	"fmt"
	"net/http"

	logpkg "github.com/sparkymat/pover/log"
)

type POVService interface {
	Compile(ctx context.Context, rubyCode string) (string, error)
}

func GenerateImage(p POVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := logpkg.FromContext(r.Context())

		err := r.ParseForm()
		if err != nil {
			log.Errorf("error parsing form: %v", err)

			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("bad request"))

			return
		}

		code := r.Form.Get("code")

		imageFilename, err := p.Compile(r.Context(), code)
		if err != nil {
			log.Errorf("error compiling code: %v", err)

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		response := fmt.Sprintf("<img src='/images/%s' />", imageFilename)

		_, _ = w.Write([]byte(response))
	}
}
