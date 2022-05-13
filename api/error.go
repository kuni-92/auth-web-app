package api

import "net/http"

func Error(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "error", nil)
}
