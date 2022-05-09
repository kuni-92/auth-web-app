package api

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	name := r.PostFormValue("name")
	password := r.PostFormValue("password")

	if password == "" || name == "" {
		return
	}

	renderTemplate(w, "top", name)
}
