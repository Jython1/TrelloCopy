package repository

import (
	"errors"
	"sync"
	"time"
	"trellocopy/internal/entity"
)

type InMemoryColumnRepository struct {
	columns map[int]*entity.Column
	counter int
	mtx     sync.RWMutex
}

func NewInMemoryColumnRepository() *InMemoryColumnRepository {
	return &InMemoryColumnRepository{
		columns: make(map[int]*entity.Column),
	}
}

func (r *InMemoryColumnRepository) Create(column *entity.Column) error {

	r.mtx.Lock()
	defer r.mtx.Unlock()

	if column == nil {
		return errors.New("Board is nill")
	}

	r.counter++
	column.ID = r.counter

	r.columns[column.ID] = column

	return nil
}

func (r *InMemoryColumnRepository) GetByID(id int) (*entity.Column, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	column, exist := r.columns[id]
	if !exist {
		return nil, errors.New("Board not found")
	}
	return column, nil
}

func (r *InMemoryColumnRepository) Update(column *entity.Column) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, exist := r.columns[column.ID]; !exist {
		return errors.New("Board is not exist")
	}
	column.UpdatedAt = time.Now()

	r.columns[column.ID] = column
	return nil
}

func (r *InMemoryColumnRepository) Delete(id int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	_, exist := r.columns[id]
	if !exist {
		return errors.New("Board not found")
	}
	delete(r.columns, id)
	return nil
}
