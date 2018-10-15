package main

import (
	"html/template"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func direct01Handler(w http.ResponseWriter, r *http.Request) {
	p := PageVariables{}
	p.Title = "Direct - 01"

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
	ExecuteTemplate("tmpl/direct.html", w, p)
}

func direct02Handler(w http.ResponseWriter, r *http.Request) {
	p := PageVariables{}
	p.Title = "Direct - 02"

	blackList := []string{
		"localhost",
		"127.0.0.1",
	}

	url := r.URL.Query().Get("url")
	if url == "" {
		p.Data = template.HTML(URLInput)
	} else {
		bad := false
		for _, u := range blackList {
			if strings.Contains(url, u) {
				bad = true
			}
		}

		if !bad {
			resp, err := SimpleGET(url)
			if err != nil {
				log.WithError(err).Warn("Could not make request")
				p.Data = template.HTML(`Unexpected error occured. Could not make request`)
			} else {
				data := string(resp)
				p.Data = template.HTML(strings.Replace(data, "\n", "<br>", -1))
			}
		} else {
			p.Data = "Invalid URL Entered. Please check you query again."
		}
	}
	ExecuteTemplate("tmpl/direct.html", w, p)
}
