package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthcheck", healthCheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Health ok")
	log.Printf("Healthcheck endpoint hit\n")
}
