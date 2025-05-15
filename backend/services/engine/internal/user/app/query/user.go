package query

import (
	"context"
	"errors"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/model"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/repository"
)

type GetUser struct {
	UserID uint
}

type GetUserHandler struct {
	repo repository.Repository
}

func NewGetUserHandler(repo repository.Repository) GetUserHandler {
	return GetUserHandler{repo: repo}
}

func (h *GetUserHandler) Handle(ctx context.Context, query GetUser) (*model.User, error) {
	if query.UserID == 0 {
		return nil, errors.New("user ID is required")
	}
	// user, err := h.repo.GetByUID(ctx, query.UserID)
	user, err := h.repo.GetByID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
