package repository

import "trellocopy/internal/entity"

type ColumnRepository interface {
	Create(column *entity.Column) error
	GetByID(id int) (*entity.Column, error)
	Update(column *entity.Column) error
	Delete(id int) error
}
