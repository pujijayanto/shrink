package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("POST /", app.shrink)
	mux.HandleFunc("GET /{slug}", app.redirectTo)

	return app.logRequest(commonHeaders(mux))
}
