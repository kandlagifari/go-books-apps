package models

import (
	"database/sql"
	"encoding/json"
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

type CustomBook struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       int       `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(CustomBook{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		ImageURL:    b.ImageURL,
		ReleaseYear: b.ReleaseYear,
		Price:       b.Price,
		TotalPage:   b.TotalPage,
		Thickness:   b.Thickness,
		CategoryID:  b.CategoryID,
		CreatedAt:   b.CreatedAt,
		CreatedBy:   b.CreatedBy.String,
		ModifiedAt:  b.ModifiedAt,
		ModifiedBy:  b.ModifiedBy.String,
	})
}
