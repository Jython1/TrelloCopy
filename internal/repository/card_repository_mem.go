package repository

import (
	"errors"
	"sync"
	"time"
	"trellocopy/internal/entity"
)

type InMemoryCardRepository struct {
	mtx     sync.RWMutex
	cards   map[int]*entity.Card
	counter int
}

func NewInMemoryCardRepository() *InMemoryCardRepository {
	return &InMemoryCardRepository{
		cards: make(map[int]*entity.Card),
	}
}

func (r *InMemoryCardRepository) Create(card *entity.Card) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if card == nil {
		return errors.New("Card is nil")
	}

	r.counter++
	card.ID = r.counter
	card.CreatedAt = time.Now()
	card.UpdatedAt = time.Now()
	r.cards[card.ID] = card

	return nil
}

func (r *InMemoryCardRepository) GetByID(id int) (*entity.Card, error) {
	r.mtx.RLock()
	defer r.mtx.Unlock()

	card, exist := r.cards[id]
	if !exist {
		return nil, errors.New("Card not found")

	}
	return card, nil
}

func (r *InMemoryCardRepository) Update(card *entity.Card) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, exist := r.cards[card.ID]; !exist {
		return errors.New("Card doesn't exist")
	}
	card.UpdatedAt = time.Now()
	// ВАЖНО: не сохраняет обновленную карточку в мапу!
	return nil
}

func (r *InMemoryCardRepository) Delete(id int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	_, exist := r.cards[id]
	if !exist {
		return errors.New("Card not found")
	}

	delete(r.cards, id)

	return nil

}
