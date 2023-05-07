package storage

import (
	"bot/internal/entity"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

// Storage for items.
type Storage struct {
	sync.RWMutex
	items      map[string]entity.Item
	allRates   int
	countRates int
}

// ErrItemNotFound error for item not found.
var ErrItemNotFound = errors.New("item not found")

var defaultItems = map[string]entity.Item{
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
}

// New returns new storage
func New() (*Storage, error) {
	f, err := os.Open("items.json")
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	defer f.Close()

	var items = make([]entity.Item, 0)

	err = json.NewDecoder(f).Decode(&items)
	if err != nil {
		return nil, fmt.Errorf("read json: %w", err)
	}

	itemsMap := make(map[string]entity.Item)
	for _, item := range items {
		itemsMap[item.Name] = item
	}

	return &Storage{
		items: itemsMap,
	}, nil
}

// NewDefault returns new storage with default items
func NewDefault() *Storage {
	return &Storage{
		items: defaultItems,
	}
}

// GetItemByName returns item by name.
func (s *Storage) GetItemByName(name string) (entity.Item, error) {
	s.RLock()
	defer s.RUnlock()

	item, ok := s.items[name]
	if !ok {
		return entity.Item{}, ErrItemNotFound
	}

	return item, nil
}

// GetAll returns the names of all items from the repository.
func (s *Storage) GetAll() []string {
	s.RLock()
	defer s.RUnlock()

	var items []string
	for key := range s.items {
		items = append(items, key)
	}
	return items
}

// GetAVG returns the average rating of the store.
func (s *Storage) GetAVG() float64 {
	s.RLock()
	defer s.RUnlock()

	return float64(s.allRates) / float64(s.countRates)
}

// AddRate adds rate to the store.
func (s *Storage) AddRate(rate int) {
	s.Lock()
	defer s.Unlock()

	s.allRates += rate
	s.countRates++
}
