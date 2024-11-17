package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID         int            `json:"id"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	CreatedAt  time.Time      `json:"created_at"`
	CreatedBy  sql.NullString `json:"created_by"`
	ModifiedAt time.Time      `json:"modified_at"`
	ModifiedBy sql.NullString `json:"modified_by"`
}
