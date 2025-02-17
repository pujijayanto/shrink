package shrinker

import (
	"crypto/tls"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/pujijayanto/shrink/internal/assert"
)

func TestValidUrl(t *testing.T) {
	tests := []struct {
		url      string
		expected bool
	}{
		{"http://example.com", true},
		{"https://example.com", true},
		{"ftp://example.com", false},
		{"invalid-url", false},
		{"", false},
	}

	for _, test := range tests {
		result := ValidUrl(test.url)
		assert.Equal(t, result, test.expected)
	}
}

func TestBuildShortUrl(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	tests := []struct {
		slug     string
		expected string
	}{
		{"abc123", "http://example.com/abc123"},
		{"xyz789", "http://example.com/xyz789"},
	}

	for _, test := range tests {
		result := BuildShortUrl(test.slug, req)
		assert.Equal(t, result, test.expected)
	}

	// Test with HTTPS
	reqTLS := httptest.NewRequest("GET", "https://example.com", nil)
	reqTLS.TLS = &tls.ConnectionState{} // Set TLS to a non-nil value to simulate HTTPS
	for _, test := range tests {
		result := BuildShortUrl(test.slug, reqTLS)
		expected := "https://example.com/" + test.slug
		assert.Equal(t, result, expected)
	}
}

func TestBuildSlug(t *testing.T) {
	text := "example text"
	slug1 := BuildSlug(text)
	time.Sleep(2 * time.Nanosecond) // simulate different timestamps
	slug2 := BuildSlug(text)

	if slug1 == slug2 {
		t.Errorf("BuildSlug() produced the same slug for different timestamps: %s", slug1)
	}

	if len(slug1) != 7 {
		t.Errorf("BuildSlug() produced a slug of incorrect length: %s", slug1)
	}
}
