package repository

import (
	"sync"
	"testing"
	"trellocopy/internal/entity"
)

func TestInMemoryColumnRepository_Create(t *testing.T) {
	repo := NewInMemoryColumnRepository()
	column := &entity.Column{Title: "Test card"}
	err := repo.Create(column)

	if err != nil {
		t.Fatal("Couldn't create column")
	}

	if column.ID == 0 {
		t.Error("Column was not set")
	}

}

func TestInMemoryColumnRepository_Delete(t *testing.T) {
	repo := NewInMemoryColumnRepository()
	column := &entity.Column{Title: "Test card"}
	err := repo.Create(column)

	if err != nil {
		t.Fatal("Couldn't create column")
	}

	err = repo.Delete(column.ID)
	if err != nil {
		t.Error("Could not delete column")
	}

	_, err = repo.GetByID(column.ID)
	if err == nil {
		t.Error("Column still exist")
	}
}

func TestInMemoryColumnRepository_Update(t *testing.T) {
	repo := NewInMemoryColumnRepository()
	column := &entity.Column{Title: "Test card"}
	err := repo.Create(column)

	if err != nil {
		t.Fatal("Couldn't create column")
	}

	column.Title = "New test card"

	err = repo.Update(column)
	if err != nil {
		t.Error("Could not update column")
	}

	updated, err := repo.GetByID(column.ID)

	if updated.Title != column.Title {
		t.Error("Could not change name")
	}

}

func TestInMemoryColumnRepository_Concurency(t *testing.T) {

	repo := NewInMemoryColumnRepository()
	column := &entity.Column{Title: "Test card"}
	err := repo.Create(column)

	if err != nil {
		t.Fatal("Couldn't create column")
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := repo.GetByID(column.ID)
			if err != nil {
				t.Errorf("Concurrent read failed: %v", err)
			}

		}()
		wg.Wait()
	}
	t.Log("1000 concurrent reads completed")

}
