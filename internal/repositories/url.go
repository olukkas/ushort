package repositories

import (
	"crypto/rand"
	"database/sql"
	"math/big"
	"time"
)

type URL struct {
	ID          int64
	Slug        string
	OriginalUrl string
	CreatedAt   time.Time
	ExpireAt    *time.Time
}

//goland:noinspection SpellCheckingInspection
var slugChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func NewURL(originalUrl string) *URL {
	slug, _ := generateSlug(10)
	return &URL{Slug: slug, OriginalUrl: originalUrl, CreatedAt: time.Now()}
}

func generateSlug(n int) (string, error) {
	b := make([]rune, n)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(slugChars))))
		if err != nil {
			return "", err
		}
		b[i] = slugChars[num.Int64()]
	}

	return string(b), nil
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

	if model.ExpireAt == nil {
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
