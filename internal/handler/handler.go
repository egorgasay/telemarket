package handler

import (
	"bot/internal/entity"
	"context"
	api "github.com/egorgasay/telemarket-grpc/telemarket"
)

type useCase interface {
	UpsertItem(ctx context.Context, i entity.IItem) error
	GetItems() ([]entity.IItem, error)
	GetItem(id string) (entity.IItem, error)
	Pause()
	Awake()
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

func (h *Handler) GetItem(ctx context.Context, r *api.GetItemRequest) (*api.GetItemResponse, error) {
	item, err := h.logic.GetItem(r.GetItemID())
	if err != nil {
		return &api.GetItemResponse{}, err
	}

	return &api.GetItemResponse{Item: &api.Item{
		Name:        item.GetName(),
		Id:          item.GetId(),
		Description: item.GetDescription(),
		Price:       item.GetPrice(),
		Image:       item.GetImage(),
		Quantity:    item.GetQuantity(),
	}}, nil
}

func (h *Handler) Pause(ctx context.Context, r *api.PauseRequest) (*api.PauseResponse, error) {
	h.logic.Pause()
	return &api.PauseResponse{}, nil
}

func (h *Handler) Awake(ctx context.Context, r *api.AwakeRequest) (*api.AwakeResponse, error) {
	h.logic.Awake()
	return &api.AwakeResponse{}, nil
}
