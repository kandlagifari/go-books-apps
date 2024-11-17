package models

import (
	"database/sql"
	"time"
)

type Book struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	ImageURL    string         `json:"image_url"`
	ReleaseYear int            `json:"release_year"`
	Price       int            `json:"price"`
	TotalPage   int            `json:"total_page"`
	Thickness   string         `json:"thickness"`
	CategoryID  int            `json:"category_id"`
	CreatedAt   time.Time      `json:"created_at"`
	CreatedBy   sql.NullString `json:"created_by"`
	ModifiedAt  time.Time      `json:"modified_at"`
	ModifiedBy  sql.NullString `json:"modified_by"`
}