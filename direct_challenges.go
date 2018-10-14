package main

import (
	"html/template"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func direct01Handler(w http.ResponseWriter, r *http.Request) {
	p := PageVariables{}
	url := r.URL.Query().Get("url")

	if url == "" {
		p.Data = template.HTML(URLInput)
	} else {
		resp, err := SimpleGET(url)
		if err != nil {
			log.WithError(err).Warn("Could not make request")
			p.Data = template.HTML(`Unexpected error occured. Could not make request`)
		} else {
			data := string(resp)
			p.Data = template.HTML(strings.Replace(data, "\n", "<br>", -1))
		}
	}
	ExecuteTemplate("tmpl/direct-01.html", w, p)
}
