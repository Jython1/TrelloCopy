package pg_repository

import (
	"database/sql"
	"trellocopy/internal/entity"
	"trellocopy/internal/repository"
)

type PGBoardRepository struct {
	db *sql.DB
}

func NewPGBoardRepository(db *sql.DB) repository.BoardRepository {
	return &PGBoardRepository{db: db}
}

func (r *PGBoardRepository) Create(board *entity.Board) error {
	query := `
        INSERT INTO boards (user_id, title, description, position)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		board.UserID,
		board.Title,
		board.Description,
		board.Position,
	).Scan(&board.ID, &board.CreatedAt, &board.UpdatedAt)

	return err
}

func (r *PGBoardRepository) Delete(id int) error {
	query := `
	DELETE FROM boards
	WHERE id = $1
	`

	result, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (r *PGBoardRepository) GetByID(id int) (*entity.Board, error) {

	query := `
        SELECT 
            id, 
            user_id, 
            position, 
            title, 
            description, 
            created_at, 
            updated_at
        FROM boards
        WHERE id = $1
    `

	board := &entity.Board{}

	err := r.db.QueryRow(query, id).Scan(
		&board.ID,
		&board.UserID,
		&board.Position,
		&board.Title,
		&board.Description,
		&board.CreatedAt,
		&board.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return board, nil
}

func (r *PGBoardRepository) Update(board entity.Board) {
	panic("unimplemented")
}
