package repository

import (
	"errors"
	"sync"
	"time"
	"trellocopy/internal/entity"
)

type InMemoryBoardRepository struct {
	boards  map[int]*entity.Board
	counter int
	mtx     sync.RWMutex
}

func NewInMemoryBoardRepository() *InMemoryBoardRepository {
	return &InMemoryBoardRepository{
		boards: make(map[int]*entity.Board),
	}
}

func (r *InMemoryBoardRepository) Create(board *entity.Board) error {

	r.mtx.Lock()
	defer r.mtx.Unlock()

	if board == nil {
		return errors.New("Board is nill")
	}

	r.counter++
	board.ID = r.counter

	r.boards[board.ID] = board

	return nil
}

func (r *InMemoryBoardRepository) GetByID(id int) (*entity.Board, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	board, exist := r.boards[id]
	if !exist {
		return nil, errors.New("Board not found")
	}
	return board, nil
}

func (r *InMemoryBoardRepository) Update(board *entity.Board) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, exist := r.boards[board.ID]; !exist {
		return errors.New("Board is not exist")
	}
	board.UpdatedAt = time.Now()

	r.boards[board.ID] = board
	return nil
}

func (r *InMemoryBoardRepository) Delete(id int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	_, exist := r.boards[id]
	if !exist {
		return errors.New("Board not found")
	}
	delete(r.boards, id)
	return nil
}
