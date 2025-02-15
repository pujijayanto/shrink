package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri), slog.String("trace", trace))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func doHashUsingSalt(text string) string {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltedText := fmt.Sprintf("text: '%s', salt: %s", text, salt)

	sha := sha256.New()
	sha.Write([]byte(saltedText))
	encrypted := sha.Sum(nil)

	// Convert to hex (will be alphanumeric only)
	hexEncoded := hex.EncodeToString(encrypted)

	// Take first 7 characters
	return hexEncoded[:7]
}

func buildShortUrl(slug string, r *http.Request) string {
	host := r.Host

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	shortUrl := fmt.Sprintf("%s://%s/%s", scheme, host, slug)
	return shortUrl
}
