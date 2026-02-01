package repository

import (
	"sync"
	"testing"
	"trellocopy/internal/entity"
)

func TestInMemoryRepo_Create(t *testing.T) {
	repo := NewInMemoryBoardRepository()

	board := &entity.Board{Title: "Test board"}
	err := repo.Create(board)
	if err != nil {
		t.Fatalf("Fail to create: %v", err)
	}

	if board.ID == 0 {
		t.Error("Board ID was not set")
	}
}

func TestInMemoryRepo_GetByID(t *testing.T) {
	repo := NewInMemoryBoardRepository()
	board := &entity.Board{Title: "Test board"}

	err := repo.Create(board)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	found, err := repo.GetByID(board.ID)

	if err != nil {
		t.Fatal(err)
	}

	if found == nil {
		t.Fatal("Found board should not be nil")
	}

	if found.Title != board.Title {
		t.Errorf("Expected name %q got %q", board.Title, found.Title)
	}

	if found.ID != board.ID {
		t.Errorf("IDs don't match: expected %d, got %d", board.ID, found.ID)
	}
}

func TestInMemoryRepo_GetByID_NotFound(t *testing.T) {
	repo := NewInMemoryBoardRepository()

	_, err := repo.GetByID(999)
	if err == nil {
		t.Error("Should return error for non-existent board")
	}
}

func TestInMemoryRepo_Delete(t *testing.T) {
	repo := NewInMemoryBoardRepository()
	board := &entity.Board{Title: "Test board"}

	err := repo.Create(board)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	err = repo.Delete(board.ID)
	if err != nil {
		t.Error("Could not delete board")
	}

	_, err = repo.GetByID(board.ID)
	if err == nil {
		t.Error("Board still exist")
	}
}

func TestInMemoryRepo_Update(t *testing.T) {
	repo := NewInMemoryBoardRepository()
	board := &entity.Board{Title: "Test board"}

	err := repo.Create(board)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	board.Title = "New test board"

	err = repo.Update(board)
	if err != nil {
		t.Error("Could not update board")
	}

	updated, err := repo.GetByID(board.ID)

	if updated.Title != board.Title {
		t.Error("Could not change name")
	}

}

func TestInMemoryRepo_Concurency(t *testing.T) {
	repo := NewInMemoryBoardRepository()

	board := &entity.Board{Title: "Test board"}
	repo.Create(board)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := repo.GetByID(board.ID)
			if err != nil {
				t.Errorf("Concurrent read failed: %v", err)
			}

		}()
		wg.Wait()
	}
	t.Log("1000 concurrent reads completed")

}
