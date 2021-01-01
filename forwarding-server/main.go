package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

const httpPort = ":80"

var apiclient = http.DefaultClient

func main() {
	mux := http.NewServeMux()
	mux.Handle("/1/", handler(handleAPIRequest))

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("cert-cache"),
		HostPolicy: autocert.HostWhitelist("testing.detachapp.com"),
	}

	server := &http.Server{
		Addr:    ":443",
		Handler: mux,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":80", certManager.HTTPHandler(mux))
	server.ListenAndServeTLS("", "")
}

func getBestApiServerHost() string {
	APIServerURLSs := []string{"api"}
	// APIServerURLSs := []string{"127.0.0.1:8083"}
	// APIServerURLSs := []string{"host.docker.internal:8083"}
	return APIServerURLSs[0]
}

func handleAPIRequest(w http.ResponseWriter, r *http.Request) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	// you can reassign the body if you need to parse it as multipart
	r.Body = ioutil.NopCloser(bytes.NewReader(body))

	r.RequestURI = strings.TrimPrefix(r.RequestURI, "/1")
	bestAPIHost := getBestApiServerHost()
	fmt.Printf("Forwarding Request %v to  %v\n", r.RequestURI, bestAPIHost)

	// create a new url from the raw RequestURI sent by the client
	url := fmt.Sprintf("%s://%s%s", "http", bestAPIHost, r.RequestURI)

	newReq, err := http.NewRequest(r.Method, url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("err!", err)
		return err
	}
	resp, err := apiclient.Do(newReq)
	if err != nil {
		fmt.Println("err1!", err)
		return err
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err2!", err)
		return err
	}

	w.Write(body)

	//TODO: 1. Check certificate
	//TODO: 2. Forward to API Server via docker or something
	//TODO: 3. Return values written by API Server

	return nil
}

type handler func(http.ResponseWriter, *http.Request) error

//I believe this just handles errors that come from calling the handler
// (it just prints the error to the HTML page)
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %s", err)
		log.Println(err)
	}
}
