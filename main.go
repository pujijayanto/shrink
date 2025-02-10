package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shrink is an URL shortener"))
}

func shrink(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Long URL become short"))
}

func redirectTo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("should redirect short URL to original URL"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/", shrink)
	mux.HandleFunc("/:slug", redirectTo)

	log.Println("starting server on :3200")
	err := http.ListenAndServe(":3200", mux)
	log.Fatal(err)
}
