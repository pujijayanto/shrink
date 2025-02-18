package mocks

import (
	"context"
	"time"

	"github.com/pujijayanto/shrink/internal/models"
)

var mockLink = models.Link{
	Id:          1,
	Slug:        "abcd123",
	OriginalURL: "https://example.com/",
	ClickCount:  1,
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

type LinkModel struct{}

func (m *LinkModel) Insert(ctx context.Context, og, slug string) (string, error) {
	return "abcd123", nil
}

func (m *LinkModel) Get(ctx context.Context, slug string) (string, error) {
	switch slug {
	case "abcd123":
		return mockLink.OriginalURL, nil
	default:
		return "", models.ErrNoRecord
	}
}
