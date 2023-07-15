package handler

import (
	"bot/internal/entity"
	"context"
	api "github.com/egorgasay/telemarket-grpc/telemarket"
)

type useCase interface {
	UpsertItem(ctx context.Context, i entity.IItem) error
}

type Handler struct {
	logic useCase
	api.UnimplementedTelemarketServer
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) UpsertItem(ctx context.Context, r *api.UpsertItemRequest) (*api.UpsertItemResponse, error) {
	err := h.logic.UpsertItem(ctx, r.GetItem())
	if err != nil {
		return &api.UpsertItemResponse{}, err
	}

	return &api.UpsertItemResponse{}, nil
}

func (h *Handler) UpdateButton(ctx context.Context, r *api.UpdateButtonRequest) (*api.UpdateButtonResponse, error) {
	// TODO: UPDATE BUTTON BY ID
	return &api.UpdateButtonResponse{}, nil
}

func (h *Handler) Ping(ctx context.Context, r *api.PingRequest) (*api.PingResponse, error) {
	return &api.PingResponse{}, nil
}
