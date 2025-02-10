package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("POST /{$}", shrink)
	mux.HandleFunc("GET /{slug}", redirectTo)

	log.Println("starting server on :3200")
	err := http.ListenAndServe(":3200", mux)
	log.Fatal(err)
}
