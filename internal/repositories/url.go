package repositories

import (
	"database/sql"
	"math/rand"
	"time"
)

type URL struct {
	ID          int64
	Slug        string
	OriginalUrl string
	CreatedAt   time.Time
	ExpireAt    time.Time
}

func NewURL(originalUrl string) *URL {
	slug := generateSlug(10)
	return &URL{Slug: slug, OriginalUrl: originalUrl}
}

func generateSlug(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (r *UrlRepository) Save(model *URL) (*URL, error) {
	var err error
	var result sql.Result

	if model.ExpireAt.IsZero() {
		result, err = r.db.Exec(
			"INSERT INTO urls (slug, original_url) VALUES (?, ?)",
			model.Slug, model.OriginalUrl,
		)
	} else {
		result, err = r.db.Exec(
			"INSERT INTO urls (slug, original_url, expires_at) VALUES (?, ?, ?)",
			model.Slug, model.OriginalUrl, model.ExpireAt,
		)
	}

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var url URL
	query := `SELECT id, slug, original_url, created_at, expires_at
	          FROM urls WHERE id = ?`
	err = r.db.QueryRow(query, id).
		Scan(&url.ID, &url.Slug, &url.OriginalUrl, &url.CreatedAt, &url.ExpireAt)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
