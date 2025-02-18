package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/pujijayanto/shrink/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("POST /", app.shrink)
	mux.HandleFunc("GET /{slug}", app.redirectTo)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
