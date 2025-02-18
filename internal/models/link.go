package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Link struct {
	Id          int64     `json:"id" db:"id"`
	Slug        string    `json:"slug" db:"slug"`
	OriginalURL string    `json:"original_url" db:"original_url"`
	ClickCount  int       `json:"click_count" db:"click_count"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type LinkModel struct {
	DB *pgxpool.Pool
}

type LinkModelInterface interface {
	Insert(ctx context.Context, originalUrl, slug string) (string, error)
	Get(ctx context.Context, slug string) (string, error)
}

func (m *LinkModel) Insert(ctx context.Context, originalUrl, slug string) (string, error) {
	now := time.Now()

	query := `
		INSERT INTO links (slug, original_url, click_count, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $4)
		RETURNING slug
	`

	var returnedSlug string
	err := m.DB.QueryRow(ctx, query, slug, originalUrl, 0, now).Scan(&returnedSlug)
	if err != nil {
		return "", err
	}

	return returnedSlug, nil
}

func (m *LinkModel) Get(ctx context.Context, slug string) (string, error) {
	var originalURL string

	tx, err := m.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	query := `
			UPDATE links
			SET click_count = click_count + 1,
					updated_at = NOW()
			WHERE slug = $1
			RETURNING original_url`

	err = tx.QueryRow(
		ctx,
		query,
		slug,
	).Scan(&originalURL)

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", ErrNoRecord
		}
		return "", err
	}

	if err = tx.Commit(ctx); err != nil {
		return "", err
	}

	return originalURL, nil
}
