package handler

import (
	"bot/internal/entity"
	"context"
	api "github.com/egorgasay/telemarket-grpc/telemarket"
)

type useCase interface {
	UpsertItem(ctx context.Context, i entity.IItem) error
	GetItems() ([]entity.IItem, error)
}

type Handler struct {
	logic useCase
	api.TelemarketServer
}

func New(logic useCase) *Handler {

	return &Handler{
		logic: logic,
	}
}

func (h *Handler) UpsertItem(ctx context.Context, r *api.UpsertItemRequest) (*api.UpsertItemResponse, error) {
	err := h.logic.UpsertItem(ctx, r.GetItem())
	if err != nil {
		return &api.UpsertItemResponse{}, err
	}

	return &api.UpsertItemResponse{}, nil
}

func (h *Handler) UpdateButton(ctx context.Context, r *api.UpdateButtonRequest) (*api.UpdateButtonResponse, error) {
	// TODO: UPDATE  BUTTON BY ID
	return &api.UpdateButtonResponse{}, nil
}

func (h *Handler) Ping(ctx context.Context, r *api.PingRequest) (*api.PingResponse, error) {
	return &api.PingResponse{}, nil
}

func (h *Handler) GetItems(ctx context.Context, r *api.GetItemsRequest) (*api.GetItemsResponse, error) {
	items, err := h.logic.GetItems()
	if err != nil {
		return &api.GetItemsResponse{}, err
	}

	var itemsWithID = make([]*api.ItemNameWithID, 0, len(items))
	for _, item := range items {
		itemsWithID = append(itemsWithID, &api.ItemNameWithID{
			Name: item.GetName(),
			Id:   item.GetId(),
		})
	}

	return &api.GetItemsResponse{Items: itemsWithID}, nil
}
