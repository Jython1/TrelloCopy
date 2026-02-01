package entity

import "time"

type Card struct {
	ID          int       `db:"id" json:"id"`
	ColumnID    int       `db:"col_id" json:"col_id"`
	Position    int       `db:"position" json:"position"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
