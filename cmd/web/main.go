package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("POST /{$}", shrink)
	mux.HandleFunc("GET /{slug}", redirectTo)

	log.Println("starting server on :3200")
	err := http.ListenAndServe(":3200", mux)
	log.Fatal(err)
}
