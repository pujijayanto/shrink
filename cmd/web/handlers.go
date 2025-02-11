package main

import (
	"log/slog"
	"net/http"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	templateFiles := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (app *application) shrink(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.logger.Info("received request", slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
	w.Write([]byte("Long URL become short"))
}

func (app *application) redirectTo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.logger.Info("received request", slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
	w.Write([]byte("should redirect short URL to original URL"))
}
