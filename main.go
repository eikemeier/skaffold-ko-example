package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)

	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
