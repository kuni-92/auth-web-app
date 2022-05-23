package api

import (
	"net/http"
)

func Top(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	renderTemplate(w, "top", name)
}
