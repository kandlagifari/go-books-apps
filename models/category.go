package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Category struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	CreatedAt  time.Time      `json:"created_at"`
	CreatedBy  sql.NullString `json:"created_by"`
	ModifiedAt time.Time      `json:"modified_at"`
	ModifiedBy sql.NullString `json:"modified_by"`
}

type CustomCategory struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func (c *Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(CustomCategory{
		ID:         c.ID,
		Name:       c.Name,
		CreatedAt:  c.CreatedAt,
		CreatedBy:  c.CreatedBy.String,
		ModifiedAt: c.ModifiedAt,
		ModifiedBy: c.ModifiedBy.String,
	})
}
