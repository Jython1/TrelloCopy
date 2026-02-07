package pg_repository

import (
	"database/sql"
	"trellocopy/internal/entity"
	"trellocopy/internal/repository"
)

type PGColumnRepository struct {
	db *sql.DB
}

func NewPGColumnRepository(db *sql.DB) repository.ColumnRepository {
	return &PGColumnRepository{db: db}
}

func (r *PGColumnRepository) Create(column *entity.Column) error {
	query := `
        INSERT INTO columns (board_id, title, description, position)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		column.BoardID,
		column.Title,
		column.Description,
		column.Position,
	).Scan(&column.ID, &column.CreatedAt, &column.UpdatedAt)

	return err
}

func (r *PGColumnRepository) Delete(id int) error {
	query := `
	DELETE FROM columns
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

func (r *PGColumnRepository) GetByID(id int) (*entity.Column, error) {

	query := `
        SELECT 
            id, 
			board_id,
            position, 
            title, 
            description, 
            created_at, 
            updated_at
        FROM columns
        WHERE id = $1
    `

	column := &entity.Column{}

	err := r.db.QueryRow(query, id).Scan(
		&column.ID,
		&column.BoardID,
		&column.Position,
		&column.Title,
		&column.Description,
		&column.CreatedAt,
		&column.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return column, nil
}

func (r *PGColumnRepository) Update(column *entity.Column) error {
	query := `
        UPDATE columns 
        SET 
            title = $1,
            description = $2,
            position = $3,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $4
        RETURNING updated_at`

	err := r.db.QueryRow(
		query,
		column.Title,
		column.Description,
		column.Position,
		column.ID,
	).Scan(&column.UpdatedAt)

	return err
}
