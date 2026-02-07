package repository

import (
	"trellocopy/internal/entity"
)

type BoardRepository interface {
	Create(board *entity.Board) error
	GetByID(id int) (*entity.Board, error)
	Update(board *entity.Board) error
	Delete(id int) error
}
