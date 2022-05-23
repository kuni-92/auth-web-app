package middleware

import (
	"log"
	"net/http"

	"github.com/kunihiro-dev/auth-web-app/api"
	"github.com/kunihiro-dev/auth-web-app/session"
)

func Auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(api.SESSION_KEY)
		if err != nil {
			log.Println("Session data not found in cookie.")
			api.Error(w, r)
			return
		} 

		s, err := session.Take(r.PostFormValue("name"))
		if err != nil {
			log.Println("Session data not found in cache data.")
			log.Println(s)
			api.Error(w, r)
			return
		}

		if c.Value != s {
			log.Println("Invalid session data.")
			api.Error(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
