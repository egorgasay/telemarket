package usecase

import (
	"bot/internal/entity"
	"bot/internal/storage"
	"errors"
)

type UseCase struct {
	storage *storage.Storage
}

var ErrEmpty = errors.New("в магазине еще нет добавленных товаров")

func New(storage *storage.Storage) *UseCase {
	return &UseCase{
		storage: storage,
	}
}

func (u *UseCase) GetItemByName(name string) (entity.Item, error) {
	return u.storage.GetItemByName(name)
}

func (u *UseCase) GetAll() ([]string, error) {
	items := u.storage.GetAll()
	if items == nil {
		return nil, ErrEmpty
	}
	return items, nil
}
