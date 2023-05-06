package storage

import (
	"bot/internal/entity"
	"errors"
)

type Storage struct {
	items map[string]entity.Item
}

var ErrItemNotFound = errors.New("item not found")

func New() *Storage {
	return &Storage{
		items: map[string]entity.Item{
			"IH8YOU WHITE ⬜️": entity.Item{},
			"IH8YOU BLACK ⬛️": entity.Item{},
		},
	}
}

func (s *Storage) GetItemByName(name string) (entity.Item, error) {
	item, ok := s.items[name]
	if !ok {
		return entity.Item{}, ErrItemNotFound
	}

	return item, nil
}

func (s *Storage) GetAll() []string {
	var items []string
	for key := range s.items {
		items = append(items, key)
	}
	return items
}
