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
	items      map[string]entity.IItem
	allRates   int
	countRates int
}

// ErrItemNotFound error for item not found.
var ErrItemNotFound = errors.New("item not found")

var defaultItems = map[string]entity.IItem{
	"1": entity.Item{
		ID:          "1",
		Name:        "HATE ⬜",
		Price:       "1500",
		Quantity:    0,
		Description: "100% хлопок.",
	},
	"2": entity.Item{
		ID:          "2",
		Name:        "HATE ⬛️",
		Price:       "1500",
		Quantity:    0,
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

	iitems := make(map[string]entity.IItem, len(items))
	for i, item := range items {
		item.ID = fmt.Sprintf("%d", i+1)
		iitems[item.ID] = item
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

	if _, ok := s.items[item.GetId()]; ok {
		s.items[item.GetId()] = item
		return nil
	}

	s.items[item.GetId()] = item

	return nil
}

func (s *Storage) GetItems() []entity.IItem {
	s.RLock()
	defer s.RUnlock()

	var items []entity.IItem
	for _, item := range s.items {
		items = append(items, item)
	}

	return items
}

func (s *Storage) GetItem(id string) (entity.IItem, error) {
	s.RLock()
	defer s.RUnlock()

	if i, ok := s.items[id]; ok {
		return i, nil
	}

	return nil, ErrItemNotFound
}
