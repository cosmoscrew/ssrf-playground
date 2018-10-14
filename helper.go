package main

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// ExecuteTemplate executes the template and sends output
func ExecuteTemplate(templ string, w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles(templ)
	if err != nil {
		log.WithError(err).Warning("Could not parse template")
	}

	err = t.Execute(w, data)
	if err != nil {
		log.WithError(err).Warning("Could not execute template")
	}
}

// PageVariables contains default variables for a template
type PageVariables struct {
	Data string
}
