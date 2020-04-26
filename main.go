package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/redact", RedactHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var response = `{"status": "2xx"}`
	io.WriteString(w, response)
}

func RedactHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "application/json")
	/* Start redaction flow from here */
	var response = `Start redaction flow from here`
	io.WriteString(w, response)
}
