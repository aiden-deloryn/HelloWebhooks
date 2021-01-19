package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleRequest)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":4567", nil); err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/payload" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		hash := sha256.Sum256([]byte("123456"))
		fmt.Print(time.Now().Local().Clock())
		fmt.Printf(" :: Host: %v, Method: %v, URL: %v\n", r.Host, r.Method, r.URL)
		fmt.Printf("Secret: %v\n", r.Header.Get("X-Hub-Signature-256"))
		fmt.Printf("Expected secret: %x\n\n", hash)
		fmt.Fprintf(w, "Success\n")
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}
