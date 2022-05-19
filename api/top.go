package api

import (
	"net/http"

	"github.com/kunihiro-dev/auth-web-app/session"
)

func Top(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")

	if _, err := session.GetSession(name); err != nil {
		Error(w, r)
		return
	}
	renderTemplate(w, "top", name)
}
