package entity

import (
	"database/sql"
	"time"
)

type Column struct {
	ID          int            `db:"id" json:"id"`
	BoardID     int            `db:"board_id" json:"board_id"`
	Position    int            `db:"position" json:"position"`
	Title       string         `db:"title" json:"title"`
	Description sql.NullString `db:"description" json:"description"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
}
