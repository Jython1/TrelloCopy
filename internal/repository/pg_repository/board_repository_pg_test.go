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

	// НАЧАЛО ИСПРАВЛЕНИЯ: Очистим таблицу перед тестом
	_, err := db.Exec("DELETE FROM boards")
	if err != nil {
		t.Fatalf("Failed to clean up before test: %v", err)
	}
	// КОНЕЦ ИСПРАВЛЕНИЯ

	repo := NewPGBoardRepository(db)

	board := &entity.Board{
		Title:    "Test board",
		Position: 1,
		UserID:   1, // ДОБАВЬТЕ ЭТО - обязательно для вашей таблицы!
	}

	err = repo.Create(board)
	if err != nil {
		t.Errorf("TEST CREATE: Could not create board: %v", err)
	}

	if board.ID == 0 {
		t.Error("TEST CREATE: Board ID should be set")
	}

	exist, err := repo.GetByID(board.ID)
	if err != nil {
		t.Fatalf("TEST GETBYID: Could not get board: %v", err)
	}

	if exist == nil {
		t.Fatal("TEST GETBYID: Board does not exist")
	}

	// НАЧАЛО ИСПРАВЛЕНИЯ: Добавьте проверку user_id
	if exist.UserID != board.UserID {
		t.Errorf("TEST GETBYID: Expected user_id %d, got %d", board.UserID, exist.UserID)
	}
	// КОНЕЦ ИСПРАВЛЕНИЯ

	err = repo.Delete(board.ID)
	if err != nil {
		t.Errorf("TEST DELETE: Could not delete board: %v", err)
	}

	// НАЧАЛО ИСПРАВЛЕНИЯ: Проверьте, что действительно удалилось
	deleted, err := repo.GetByID(board.ID)
	if err == nil && deleted != nil {
		t.Error("TEST DELETE: Board should be deleted but still exists")
	}
	// КОНЕЦ ИСПРАВЛЕНИЯ
}

func ConnectDB(t *testing.T) *sql.DB {
	connectionString := "postgres://jython:123@localhost:5432/trello_db?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatalf("Error openning DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatalf("Could not connect to DB: %v", err)
	}

	return db
}
