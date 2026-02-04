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

	board := &entity.Board{
		Title:    "Test board",
		Position: 1,
	}

	err := repo.Create(board)

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

	err = repo.Delete(board.ID)

	if err != nil {
		t.Errorf("TEST DELETE: Could not delete board: %v", err)
	}

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

/*
# 1. Убедись, что БД запущена
docker-compose up postgres -d

# 2. Запусти тест
go test ./internal/repository/pg_repository/... -v -run TestPGBoardRepository_GetByID

*/
