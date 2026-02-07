package pg_repository

import (
	"database/sql"
	"testing"
	"trellocopy/internal/entity"

	_ "github.com/lib/pq"
)

func TestBoardCRUD(t *testing.T) {
	db := ConnectDB(t)
	defer db.Close()

	repo := NewPGBoardRepository(db)

	testUserID := 1

	board := &entity.Board{
		Title:    "Test board",
		Position: 1,
		UserID:   testUserID,
	}

	// CREATE
	err := repo.Create(board)
	if err != nil {
		t.Fatalf("CREATE failed: %v", err)
	}
	if board.ID == 0 {
		t.Fatal("CREATE: Board ID should be set")
	}

	// Clear after test
	defer func() {
		_ = repo.Delete(board.ID)
	}()

	// GET BY ID
	fetched, err := repo.GetByID(board.ID)
	if err != nil {
		t.Fatalf("GET BY ID failed: %v", err)
	}
	if fetched == nil {
		t.Fatal("GET BY ID: Board should exist")
	}
	if fetched.Title != board.Title {
		t.Errorf("GET BY ID: Expected title %q, got %q", board.Title, fetched.Title)
	}

	// UPDATE
	newTitle := "Updated Board Title"
	board.Title = newTitle
	board.Position = 2

	err = repo.Update(board)
	if err != nil {
		t.Fatalf("UPDATE failed: %v", err)
	}

	updated, err := repo.GetByID(board.ID)
	if err != nil {
		t.Fatalf("GET AFTER UPDATE failed: %v", err)
	}
	if updated.Title != newTitle {
		t.Errorf("UPDATE: Expected title %q, got %q", newTitle, updated.Title)
	}
	if updated.Position != 2 {
		t.Errorf("UPDATE: Expected position %d, got %d", 2, updated.Position)
	}

	// DELETE
	err = repo.Delete(board.ID)
	if err != nil {
		t.Fatalf("DELETE failed: %v", err)
	}

	deleted, err := repo.GetByID(board.ID)
	if err != nil {
		t.Fatalf("GET AFTER DELETE failed: %v", err)
	}
	if deleted != nil {
		t.Fatal("DELETE: Board should not exist after deletion")
	}
}

func TestColumnCRUD(t *testing.T) {
	db := ConnectDB(t)
	defer db.Close()

	boardRepo := NewPGBoardRepository(db)
	colRepo := NewPGColumnRepository(db)

	testUserID := 1

	board := &entity.Board{
		Title:    "Test board",
		Position: 1,
		UserID:   testUserID,
	}

	err := boardRepo.Create(board)
	if err != nil {
		t.Fatalf("CREATE failed: %v", err)
	}
	if board.ID == 0 {
		t.Fatal("CREATE: Board ID should be set")
	}

	column := &entity.Column{
		Title:    "Test card",
		BoardID:  board.ID,
		Position: 1,
	}

	err = colRepo.Create(column)
	if err != nil {
		t.Fatalf("CREATE failed: %v", err)
	}
	if column.ID == 0 {
		t.Fatal("CREATE: Column ID should be set")
	}

	// Clear after test
	defer func() {
		_ = boardRepo.Delete(board.ID)
		_ = colRepo.Delete(column.ID)
	}()

	// GET BY ID
	fetched, err := colRepo.GetByID(column.ID)
	if err != nil {
		t.Fatalf("GET BY ID failed: %v", err)
	}
	if fetched == nil {
		t.Fatal("GET BY ID: Column should exist")
	}
	if fetched.Title != column.Title {
		t.Errorf("GET BY ID: Expected title %q, got %q", column.Title, fetched.Title)
	}

	// UPDATE
	newTitle := "Updated Board Title"
	column.Title = newTitle
	column.Position = 2

	err = colRepo.Update(column)
	if err != nil {
		t.Fatalf("UPDATE failed: %v", err)
	}

	updated, err := colRepo.GetByID(column.ID)
	if err != nil {
		t.Fatalf("GET AFTER UPDATE failed: %v", err)
	}
	if updated.Title != newTitle {
		t.Errorf("UPDATE: Expected title %q, got %q", newTitle, updated.Title)
	}
	if updated.Position != 2 {
		t.Errorf("UPDATE: Expected position %d, got %d", 2, updated.Position)
	}

	// DELETE
	err = colRepo.Delete(column.ID)
	if err != nil {
		t.Fatalf("DELETE failed: %v", err)
	}

	deleted, err := colRepo.GetByID(column.ID)
	if err != nil {
		t.Fatalf("GET AFTER DELETE failed: %v", err)
	}
	if deleted != nil {
		t.Fatal("DELETE: Column should not exist after deletion")
	}
}

func TestCardCRUD(t *testing.T) {
	db := ConnectDB(t)
	defer db.Close()

	boardRepo := NewPGBoardRepository(db)
	colRepo := NewPGColumnRepository(db)
	cardRepo := NewPGCardRepository(db)

	testUserID := 1

	board := &entity.Board{
		Title:    "Test board",
		Position: 1,
		UserID:   testUserID,
	}

	err := boardRepo.Create(board)
	if err != nil {
		t.Fatalf("CREATE failed: %v", err)
	}
	if board.ID == 0 {
		t.Fatal("CREATE: Board ID should be set")
	}

	column := &entity.Column{
		Title:    "Test column",
		BoardID:  board.ID,
		Position: 1,
	}

	err = colRepo.Create(column)
	if err != nil {
		t.Fatalf("CREATE failed: %v", err)
	}
	if column.ID == 0 {
		t.Fatal("CREATE: Column ID should be set")
	}

	card := &entity.Card{
		Title:    "Test card",
		ColumnID: column.ID,
		Position: 1,
	}

	err = cardRepo.Create(card)
	if err != nil {
		t.Fatalf("CREATE failed: %v", err)
	}
	if card.ID == 0 {
		t.Fatal("CREATE: Card ID should be set")
	}

	// Clear after test
	defer func() {
		_ = cardRepo.Delete(card.ID)
		_ = colRepo.Delete(column.ID)
		_ = boardRepo.Delete(board.ID)
	}()

	// GET BY ID
	fetched, err := cardRepo.GetByID(card.ID)
	if err != nil {
		t.Fatalf("GET BY ID failed: %v", err)
	}
	if fetched == nil {
		t.Fatal("GET BY ID: Card should exist")
	}
	if fetched.Title != card.Title {
		t.Errorf("GET BY ID: Expected title %q, got %q", card.Title, fetched.Title)
	}

	// UPDATE
	newTitle := "Updated Card Title"
	card.Title = newTitle
	card.Position = 2

	err = cardRepo.Update(card)
	if err != nil {
		t.Fatalf("UPDATE failed: %v", err)
	}

	updated, err := cardRepo.GetByID(card.ID)
	if err != nil {
		t.Fatalf("GET AFTER UPDATE failed: %v", err)
	}
	if updated.Title != newTitle {
		t.Errorf("UPDATE: Expected title %q, got %q", newTitle, updated.Title)
	}
	if updated.Position != 2 {
		t.Errorf("UPDATE: Expected position %d, got %d", 2, updated.Position)
	}

	// DELETE
	err = cardRepo.Delete(card.ID)
	if err != nil {
		t.Fatalf("DELETE failed: %v", err)
	}

	deleted, err := cardRepo.GetByID(card.ID)
	if err != nil {
		t.Fatalf("GET AFTER DELETE failed: %v", err)
	}
	if deleted != nil {
		t.Fatal("DELETE: Card should not exist after deletion")
	}
}

func ConnectDB(t *testing.T) *sql.DB {
	connStr := "postgres://jython:123@localhost:5432/trello_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Open DB failed: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("DB ping failed: %v", err)
	}

	return db
}
