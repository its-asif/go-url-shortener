package models

import "time"

type URL struct {
	ID          int       `db:"id"`
	ShortCode   string    `db:"short_code"`
	OriginalURL string    `db:"original_url"`
	CreatedAt   time.Time `db:"created_at"`
}
