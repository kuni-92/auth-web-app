package api

import (
	"html/template"
	"io"
	"log"
)

var Templates *template.Template
var SessionID map[string]int64

func init() {
	SessionID = make(map[string]int64)
}

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

type User struct {
	name string
	password string
}
