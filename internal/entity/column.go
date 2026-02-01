package entity

import "time"

type Column struct {
	ID          int       `db:"id" json:"id"`
	Position    int       `db:"position" json:"position"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
