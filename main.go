package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	port = flag.Int("p", 8001, "Port to listen on")
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate("tmpl/index.html", w, nil)
}

func direct01Handler(w http.ResponseWriter, r *http.Request) {
	p := PageVariables{}
	url := r.URL.Query().Get("url")

	if url == "" {
		p.Data = "No \"url\" parameter specified in the query"
	} else {
		resp, err := http.Get(url)
		if err != nil {
			log.WithError(err).Warn("Could not get url")
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithError(err).Warn("Could not read body")
			return
		}
		resp.Body.Close()

		data := string(body)
		data = strings.Replace(data, "\n", "<br>", -1)
		p.Data = data
	}
	ExecuteTemplate("tmpl/direct-01.html", w, p)
}

func main() {
	log.WithFields(log.Fields{
		"port": *port,
	}).Info("Starting SSRF-Playground web server")
	log.Info("Written By - CosmosCrew")

	go listenBackend()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/direct-01", direct01Handler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
