package storage

import (
	"bot/internal/entity"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

// Storage for items.
type Storage struct {
	sync.RWMutex
	items      []entity.IItem
	allRates   int
	countRates int
}

// ErrItemNotFound error for item not found.
var ErrItemNotFound = errors.New("item not found")

var defaultItems = []entity.IItem{
	entity.Item{
		ID:          "1",
		Name:        "HATE ⬜",
		Price:       "1500",
		Quantity:    "0",
		Description: "100% хлопок.",
	},
	entity.Item{
		ID:          "2",
		Name:        "HATE ⬛️",
		Price:       "1500",
		Quantity:    "0",
		Description: "100% хлопок.",
	},
}

// New returns new storage
func New(path string) (*Storage, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	defer f.Close()

	var items = make([]entity.Item, 0)

	err = json.NewDecoder(f).Decode(&items)
	if err != nil {
		return nil, fmt.Errorf("read json: %w", err)
	}

	iitems := make([]entity.IItem, len(items))
	for _, item := range items {
		iitems = append(iitems, item)
	}

	return &Storage{
		items: iitems,
	}, nil
}

// NewDefault returns new storage with default items
func NewDefault() *Storage {
	return &Storage{
		items: defaultItems,
	}
}

// GetItemByName returns item by name.
func (s *Storage) GetItemByName(name string) (i entity.IItem, err error) {
	s.RLock()
	defer s.RUnlock()

	for _, item := range s.items {
		if item.GetName() == name {
			return item, nil
		}
	}

	return i, ErrItemNotFound
}

// GetAll returns the names of all items from the repository.
func (s *Storage) GetAll() []string {
	s.RLock()
	defer s.RUnlock()

	var items []string
	for _, item := range s.items {
		items = append(items, item.GetName())
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

func (s *Storage) UpsertItem(ctx context.Context, item entity.IItem) error {
	s.Lock()
	defer s.Unlock()

	if ctx.Err() != nil {
		return ctx.Err()
	}

	for indx, i := range s.items {
		if i.GetName() == item.GetName() {
			s.items[indx] = item
			return nil
		}
	}

	s.items = append(s.items, item)

	return nil
}
