package main

import (
	"errors"
	"net/http"
	"strings"
	"text/template"

	"github.com/pujijayanto/shrink/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	templateFiles := []string{
		"./ui/html/index.html",
	}

	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := struct {
		OriginalURL string
	}{
		OriginalURL: "",
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) shrink(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		app.logger.Info("parse form error", "error", err)
		app.clientError(w, http.StatusBadRequest)
		return
	}

	originalUrl := r.Form.Get("url")
	if originalUrl == "" {
		app.logger.Info("URL is empty")
		app.clientError(w, http.StatusBadRequest)
		return
	}

	slug := doHashUsingSalt(originalUrl)
	insertedSlug, err := app.links.Insert(r.Context(), originalUrl, slug)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	shortenedUrl := buildShortUrl(insertedSlug, r)
	w.Write([]byte(shortenedUrl))
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

	templateFiles := []string{
		"./ui/html/index.html",
	}

	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := struct {
		OriginalURL string
	}{
		OriginalURL: originalUrl,
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
