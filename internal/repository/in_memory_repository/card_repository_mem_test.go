package repository

import (
	"sync"
	"testing"
	"trellocopy/internal/entity"
)

func TestInMemoryCardRepository_Create(t *testing.T) {
	repo := NewInMemoryCardRepository()
	card := &entity.Card{Title: "Test card"}
	err := repo.Create(card)

	if err != nil {
		t.Fatal("Couldn't create card")
	}

	if card.ID == 0 {
		t.Error("Card was not set")
	}

}

func TestInMemoryCardRepository_Delete(t *testing.T) {
	repo := NewInMemoryCardRepository()
	card := &entity.Card{Title: "Test card"}
	err := repo.Create(card)

	if err != nil {
		t.Fatal("Couldn't create card")
	}

	err = repo.Delete(card.ID)
	if err != nil {
		t.Error("Could not delete card")
	}

	_, err = repo.GetByID(card.ID)
	if err == nil {
		t.Error("Card still exist")
	}
}

func TestInMemoryCardRepository_Update(t *testing.T) {
	repo := NewInMemoryCardRepository()
	card := &entity.Card{Title: "Test card"}
	err := repo.Create(card)

	if err != nil {
		t.Fatal("Couldn't create card")
	}

	card.Title = "New test card"

	err = repo.Update(card)
	if err != nil {
		t.Error("Could not update card")
	}

	updated, err := repo.GetByID(card.ID)

	if updated.Title != card.Title {
		t.Error("Could not change name")
	}

}

func TestInMemoryCardRepository_Concurency(t *testing.T) {

	repo := NewInMemoryCardRepository()
	card := &entity.Card{Title: "Test card"}
	err := repo.Create(card)

	if err != nil {
		t.Fatal("Couldn't create card")
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := repo.GetByID(card.ID)
			if err != nil {
				t.Errorf("Concurrent read failed: %v", err)
			}

		}()
		wg.Wait()
	}
	t.Log("1000 concurrent reads completed")

}
