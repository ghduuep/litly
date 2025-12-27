package domain

import "time"

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	ISBN        string    `json:"isbn"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Genre       string    `json:"genre"`
	Pages       int       `json:"pages"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
