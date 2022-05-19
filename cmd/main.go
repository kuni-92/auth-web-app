package main

import (
	"log"
	"net/http"

	"github.com/kunihiro-dev/auth-web-app/api"
	"github.com/kunihiro-dev/auth-web-app/middleware"
)

func main() {
	// load template
	if err := api.LoadTemplate(); err != nil {
		log.Fatal("load template error.")
	}

	indexHandler := http.HandlerFunc(api.Index)
	loginHandler := http.HandlerFunc(api.Login)
	topmenuHandler := http.HandlerFunc(api.Top)

	m := http.NewServeMux()
	m.Handle("/", middleware.Logger(indexHandler))
	m.Handle("/login", middleware.Logger(loginHandler))
	m.Handle("/top", middleware.Logger(topmenuHandler))

	if err := http.ListenAndServe(":3030", m); err != nil {
		log.Println("http server error.")
	}
}
