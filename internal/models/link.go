package models

import (
	"time"

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

func (m *LinkModel) Insert(originalUrl, slug string) error {
	return nil
}

func (m *LinkModel) Get(slug string) (Link, error) {
	return Link{}, nil
}
