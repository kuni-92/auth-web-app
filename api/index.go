package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/login.tmpl","tmpl/header.tmpl", "tmpl/footer.tmpl")
	if err != nil {
		log.Println("load template error.")
		fmt.Fprintf(w, "load template error.")
	}

	if err = t.Execute(w, nil); err != nil {
		log.Println("load template error.")
		fmt.Fprintf(w, "load template error.")
	}
}
