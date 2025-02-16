package shrinker

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func ValidUrl(originalUrl string) bool {
	parsedUrl, err := url.ParseRequestURI(originalUrl)
	if err != nil {
		return false
	}

	// Ensure the URL has a valid scheme (http or https)
	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return false
	}

	return true
}

func BuildShortUrl(slug string, r *http.Request) string {
	host := r.Host

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	shortUrl := fmt.Sprintf("%s://%s/%s", scheme, host, slug)
	return shortUrl
}

func BuildSlug(text string) string {
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
