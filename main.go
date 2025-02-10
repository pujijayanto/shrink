package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Shrink is an URL shortener"))
}

func shrink(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Long URL become short"))
}

func redirectTo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("should redirect short URL to original URL"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("POST /{$}", shrink)
	mux.HandleFunc("GET /{slug}", redirectTo)

	log.Println("starting server on :3200")
	err := http.ListenAndServe(":3200", mux)
	log.Fatal(err)
}
