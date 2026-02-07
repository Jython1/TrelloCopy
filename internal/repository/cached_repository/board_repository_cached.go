package cachedrepository

import (
	"trellocopy/internal/entity"
	"trellocopy/internal/repository"
	redisrepository "trellocopy/internal/repository/redis_repository"
)

type CachedBoardRepository struct {
	boardRepo repository.BoardRepository
	redisRepo *redisrepository.RedisBoardRepository
}

func NewCachedBoardRepository(boardRepo repository.BoardRepository,
	redisRepo *redisrepository.RedisBoardRepository) repository.BoardRepository {
	return &CachedBoardRepository{
		boardRepo: boardRepo,
		redisRepo: redisRepo,
	}
}

func (r *CachedBoardRepository) Create(board *entity.Board) error {
	panic("unimplemented")
}

func (r *CachedBoardRepository) Delete(id int) error {
	panic("unimplemented")
}

func (r *CachedBoardRepository) GetByID(id int) (*entity.Board, error) {
	// Try cache first
	cachedBoard, err := r.redisRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if cachedBoard != nil {
		return cachedBoard, nil
	}

	// Cache miss - get from DB
	board, err := r.boardRepo.GetByID(id)
	if err != nil || board == nil {
		return board, err
	}

	// Cache for future requests
	_ = r.redisRepo.Set(*board)
	return board, nil
}

func (r *CachedBoardRepository) Update(board *entity.Board) error {
	panic("unimplemented")
}
