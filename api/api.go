package api

import (
	"html/template"
	"io"
	"log"
)

var Templates *template.Template

func LoadTemplate() error {
	t, err := template.ParseGlob("tmpl/*.tmpl")
	Templates = t
	return err
}

func renderTemplate(w io.Writer, name string) {
	if err := Templates.ExecuteTemplate(w, name, nil); err != nil {
		log.Printf("template rendering error. page is %s\n", name)
	}
}
