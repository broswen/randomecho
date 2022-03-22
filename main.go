package main

import (
	"fmt"
	"github.com/broswen/randomecho/counter"
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

	counter, err := counter.New(os.Getenv("CACHE_ADDR"))
	if err != nil {
		log.Fatalf("new counter: %s", err)
	}

	mux := http.DefaultServeMux
	mux.HandleFunc("/time", accessLogger(timeHandler))
	mux.HandleFunc("/env", accessLogger(envHandler))
	mux.HandleFunc("/counter/incr", incrHandler(counter))
	mux.HandleFunc("/counter", counterHandler(counter))
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

func incrHandler(counter *counter.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, err := counter.Incr(r.Context())
		fmt.Fprintf(w, "%d", val)
		defer func(err error) {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}(err)
	}
}

func counterHandler(counter *counter.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, err := counter.Get(r.Context())
		if err == nil {
			fmt.Fprintf(w, "%d", val)
		}
		defer func(err error) {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}(err)
	}
}
