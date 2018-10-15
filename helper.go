package main

import (
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

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
	Data template.HTML
	Title string
}

// SimpleGET makes a GET request to a URL
func SimpleGET(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}
	resp.Body.Close()

	return body, nil
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

//RandStringBytesMaskImprSrc from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
