package main

import (
	"net/http"
	"testing"

	"github.com/pujijayanto/shrink/internal/assert"
)

func TestHome(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.get(t, "/")
	assert.Equal(t, code, http.StatusOK)
}

func TestRedirectTo(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name       string
		slug       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "slug exist",
			slug:       "abcd123",
			wantStatus: http.StatusOK,
		},
		{
			name:       "slug does not exist",
			slug:       "nonexistent",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, _ := ts.get(t, "/"+tt.slug)

			assert.Equal(t, tt.wantStatus, code)
		})
	}
}

func TestShrink(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name       string
		formData   map[string]string
		wantStatus int
		wantBody   string
	}{
		{
			name: "valid url",
			formData: map[string]string{
				"url": "http://example.com",
			},
			wantStatus: http.StatusOK,
			wantBody:   "abcd123",
		},
		{
			name: "empty url",
			formData: map[string]string{
				"url": "",
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "Bad Request",
		},
		{
			name: "invalid url",
			formData: map[string]string{
				"url": "invalid@url.com",
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "Bad Request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.postForm(t, "/", tt.formData)

			assert.Equal(t, code, tt.wantStatus)
			assert.StringContains(t, body, tt.wantBody)
		})
	}
}
