package api

import (
	"html/template"
	"io"
	"log"
)

var Templates *template.Template
const SESSION_KEY = "authappkey"

func LoadTemplate() error {
	t, err := template.ParseGlob("tmpl/*.tmpl")
	Templates = t
	return err
}

func renderTemplate(w io.Writer, name string, param interface{}) {
	if err := Templates.ExecuteTemplate(w, name, param); err != nil {
		log.Printf("template rendering error. page is %s\n", name)
	}
}
