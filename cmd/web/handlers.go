package main

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"text/template"

	"github.com/pujijayanto/shrink/internal/models"
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
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) shrink(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	originalUrl := "google.com"
	slug := "12345ab"

	insertedSlug, err := app.links.Insert(context.TODO(), originalUrl, slug)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.Write([]byte(insertedSlug))
}

func (app *application) redirectTo(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	originalUrl, err := app.links.Get(r.Context(), slug)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	if !strings.HasPrefix(originalUrl, "http://") && !strings.HasPrefix(originalUrl, "https://") {
		originalUrl = "https://" + originalUrl
	}

	http.Redirect(w, r, originalUrl, http.StatusPermanentRedirect)
}
