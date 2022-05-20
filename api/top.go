package api

import (
	"log"
	"net/http"

	"github.com/kunihiro-dev/auth-web-app/session"
)

func Top(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")

	c, err := r.Cookie(SESSION_KEY)
	if err != nil {
		log.Println("Session data not found in cookie.")
		Error(w, r)
		return
	} 

	s, err := session.Take(name)
	if err != nil {
		log.Println("Session data not found in cache data.")
		log.Println(s)
		Error(w, r)
		return
	}

	if c.Value != s {
		log.Println("Invalid session data.")
		Error(w, r)
		return
	}
	renderTemplate(w, "top", name)
}
