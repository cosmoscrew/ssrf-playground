package main

import (
	"crypto/sha1"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func listenBackend() {
	log.WithField("url", "localhost:8083").Info("Listening with backend")

	backend := http.NewServeMux()
	backend.HandleFunc("/flag", handleBackend)

	backendServer := http.Server{
		Addr:    "localhost:8083",
		Handler: backend,
	}
	log.Fatal(backendServer.ListenAndServe())
}

func handleBackend(w http.ResponseWriter, r *http.Request) {
	sha1Hash := fmt.Sprintf("%x", sha1.Sum([]byte(RandStringBytesMaskImprSrc(12))))
	fmt.Fprintf(w, "flag{w1nn3r_w1nn3r_ch1ck3n_d1nn3r_%s}", sha1Hash)
}
