package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	port := os.Getenv("NOMAD_PORT_http")
	if port == "" {
		port = "8080"
	}

	mux := http.DefaultServeMux
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		defer req.Body.Close()
		fmt.Printf("%s %s %s %s \n %s", time.Now().Format(time.RFC3339), req.RemoteAddr, req.Method, req.URL.Path, body)
		fmt.Fprintf(rw, "%s %s %s %s \n %s", time.Now().Format(time.RFC3339), req.RemoteAddr, req.Method, req.URL.Path, body)
	})
	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatal(err)
	}
}
