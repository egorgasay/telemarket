package storage

import (
	"bot/internal/entity"
	"errors"
	"sync"
)

type Storage struct {
	sync.RWMutex
	items      map[string]entity.Item
	allRates   int
	countRates int
}

var ErrItemNotFound = errors.New("item not found")

func New() *Storage {
	return &Storage{
		items: map[string]entity.Item{
			"HATE ⬜️": {
				Name:        "HATE ⬜",
				Price:       1500,
				Quantity:    0,
				Description: "100% хлопок.",
			},
			"HATE ⬛️": {
				Name:        "HATE ⬛️",
				Price:       1500,
				Quantity:    0,
				Description: "100% хлопок.",
			},
		},
	}
}

func (s *Storage) GetItemByName(name string) (entity.Item, error) {
	s.RLock()
	defer s.RUnlock()

	item, ok := s.items[name]
	if !ok {
		return entity.Item{}, ErrItemNotFound
	}

	return item, nil
}

func (s *Storage) GetAll() []string {
	s.RLock()
	defer s.RUnlock()

	var items []string
	for key := range s.items {
		items = append(items, key)
	}
	return items
}

func (s *Storage) GetAVG() float64 {
	s.RLock()
	defer s.RUnlock()

	return float64(s.allRates) / float64(s.countRates)
}

func (s *Storage) AddRate(rate int) {
	s.Lock()
	defer s.Unlock()

	s.allRates += rate
	s.countRates++
}
