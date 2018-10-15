package main

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	port = flag.Int("p", 8001, "Port to listen on")
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate("tmpl/index.html", w, nil)
}

func main() {
	log.WithFields(log.Fields{
		"port": *port,
	}).Info("Starting SSRF-Playground web server")
	log.Info("Written By - CosmosCrew")

	go listenBackend()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/direct-01", direct01Handler)
	http.HandleFunc("/direct-02", direct02Handler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
