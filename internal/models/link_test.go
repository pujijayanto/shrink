package models

import (
	"context"
	"testing"

	"github.com/pujijayanto/shrink/internal/assert"
)

func TestLinkModel_Insert(t *testing.T) {
	tests := []struct {
		name        string
		originalURL string
		slug        string
		wantSlug    string
		wantErr     bool
	}{
		{
			name:        "valid insert",
			originalURL: "http://example.com",
			slug:        "random1",
			wantSlug:    "random1",
			wantErr:     false,
		},
		{
			name:        "insert duplicate slug",
			originalURL: "http://example.com",
			slug:        "abc123", // exist in setup.sql
			wantSlug:    "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			model := &LinkModel{DB: db}

			slug, err := model.Insert(context.Background(), tt.originalURL, tt.slug)

			assert.Equal(t, slug, tt.wantSlug)
			if !tt.wantErr {
				assert.NilError(t, err)
			}
		})
	}
}

func TestLinkModel_Get(t *testing.T) {
	tests := []struct {
		name    string
		slug    string
		wantURL string
		wantErr error
	}{
		{
			name:    "get valid slug",
			slug:    "abc123",
			wantURL: "http://example.com",
			wantErr: nil,
		},
		{
			name:    "get non existing slug",
			slug:    "nonexistent",
			wantURL: "",
			wantErr: ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			model := &LinkModel{DB: db}

			url, err := model.Get(context.Background(), tt.slug)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantURL, url)
		})
	}
}
