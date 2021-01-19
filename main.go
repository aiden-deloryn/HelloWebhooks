package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
		gotHash := strings.Split(r.Header.Get("X-Hub-Signature-256"), "=")[1]
		reqBody, _ := ioutil.ReadAll(r.Body)
		hash := hmac.New(sha256.New, []byte("123456"))
		hash.Write(reqBody)
		expectedHash := hex.EncodeToString(hash.Sum(nil))

		fmt.Print(time.Now().Local().Clock())
		fmt.Printf(" :: Host: %v, Method: %v, URL: %v\n", r.Host, r.Method, r.URL)
		fmt.Printf("Got hash: %v\n", gotHash)
		fmt.Printf("Exp hash: %s\n", expectedHash)
		fmt.Printf("Hashes match? %v\n\n", hmac.Equal([]byte(gotHash), []byte(expectedHash)))

		fmt.Fprintf(w, "Success\n")
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}
