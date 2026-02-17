package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Jenkins CI/CD 🚀")
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Server started on :8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}

