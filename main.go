package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	port := os.Getenv("NOMAD_PORT_http")
	if port == "" {
		port = "8080"
	}

	mux := http.DefaultServeMux
	mux.HandleFunc("/time", accessLogger(timeHandler))
	mux.HandleFunc("/env", accessLogger(envHandler))
	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatal(err)
	}
}

func accessLogger(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL.Path)
		handlerFunc(w, req)
	}
}

func timeHandler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "%s", time.Now().Format(time.RFC3339))
	defer func(err error) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}(err)
}

func envHandler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "%s", strings.Join(os.Environ(), "\n"))
	defer func(err error) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}(err)
}
