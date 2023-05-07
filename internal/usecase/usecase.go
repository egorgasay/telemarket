package usecase

import (
	"bot/internal/entity"
	"bot/internal/storage"
	"errors"
	"math"
)

// UseCase is a logic layer for the application.
type UseCase struct {
	storage *storage.Storage
}

// ErrEmpty is returned when the storage is empty.
var ErrEmpty = errors.New("в магазине еще нет добавленных товаров")

// New returns a new UseCase.
func New(storage *storage.Storage) *UseCase {
	return &UseCase{
		storage: storage,
	}
}

// GetItemByName returns the item with the given name from the repository.
func (u *UseCase) GetItemByName(name string) (entity.Item, error) {
	return u.storage.GetItemByName(name)
}

// GetAll returns the names of all items from the repository.
func (u *UseCase) GetAll() ([]string, error) {
	items := u.storage.GetAll()
	if items == nil {
		return nil, ErrEmpty
	}
	return items, nil
}

// AddRate adds rate to storage.
func (u *UseCase) AddRate(rate int) {
	u.storage.AddRate(rate)
}

// GetRate returns the average rating of the store.
func (u *UseCase) GetRate() float64 {
	avg := math.Ceil(u.storage.GetAVG()*100) / 100
	if math.IsNaN(avg) {
		return 0
	}
	return avg
}
