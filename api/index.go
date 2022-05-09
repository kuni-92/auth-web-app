package api

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", nil)
}
