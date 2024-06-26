package handler

import (
	"context"
	"fmt"
	"net/http"
)

type POVService interface {
	Compile(ctx context.Context, rubyCode string) (string, error)
}

func GenerateImage(p POVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		code := r.Form.Get("code")

		imageFilename, err := p.Compile(r.Context(), code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		response := fmt.Sprintf("<img src='/images/%s' />", imageFilename)

		w.Write([]byte(response))
	}
}
