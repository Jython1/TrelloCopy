package entity

import (
	"database/sql"
	"time"
)

type Card struct {
	ID          int            `db:"id" json:"id"`
	ColumnID    int            `db:"column_id" json:"column_id"`
	Position    int            `db:"position" json:"position"`
	Title       string         `db:"title" json:"title"`
	Description sql.NullString `db:"description" json:"description"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
}
