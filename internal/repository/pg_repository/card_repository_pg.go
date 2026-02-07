package pg_repository

import (
	"database/sql"
	"trellocopy/internal/entity"
	"trellocopy/internal/repository"
)

type PGCardRepository struct {
	db *sql.DB
}

func NewPGCardRepository(db *sql.DB) repository.CardRepository {
	return &PGCardRepository{db: db}
}

func (r *PGCardRepository) Create(card *entity.Card) error {
	query := `
        INSERT INTO cards (column_id, title, description, position)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		card.ColumnID,
		card.Title,
		card.Description,
		card.Position,
	).Scan(&card.ID, &card.CreatedAt, &card.UpdatedAt)

	return err
}

func (r *PGCardRepository) Delete(id int) error {
	query := `
	DELETE FROM cards
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

func (r *PGCardRepository) GetByID(id int) (*entity.Card, error) {

	query := `
        SELECT 
            id, 
			column_id,
            position, 
            title, 
            description, 
            created_at, 
            updated_at
        FROM cards
        WHERE id = $1
    `

	card := &entity.Card{}

	err := r.db.QueryRow(query, id).Scan(
		&card.ID,
		&card.ColumnID,
		&card.Position,
		&card.Title,
		&card.Description,
		&card.CreatedAt,
		&card.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (r *PGCardRepository) Update(card *entity.Card) error {
	query := `
        UPDATE cards 
        SET 
            title = $1,
            description = $2,
            position = $3,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $4
        RETURNING updated_at`

	err := r.db.QueryRow(
		query,
		card.Title,
		card.Description,
		card.Position,
		card.ID,
	).Scan(&card.UpdatedAt)

	return err
}
