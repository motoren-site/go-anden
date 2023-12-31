package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Helloooi!\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Println("Listing for requests at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
