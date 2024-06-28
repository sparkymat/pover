package handler

import (
	"fmt"
	"net/http"
)

func GenerateImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf("code: %s\n", r.Form.Get("code"))

		imageFile := "fire.jpg"
		response := fmt.Sprintf("<img src='/images/%s' />", imageFile)

		w.Write([]byte(response))
	}
}
