package repository

import "trellocopy/internal/entity"

type CardRepository interface {
	Create(card *entity.Card) error
	GetByID(id int) (*entity.Card, error)
	Update(board *entity.Card)
	Delete(id int) error
}
