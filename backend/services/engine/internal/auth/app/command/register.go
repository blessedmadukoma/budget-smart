package command

import (
	"context"
	"fmt"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/types"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/decorator"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/model"
	userRepo "github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/repository"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/messages"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/password"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/validator"
	"github.com/sirupsen/logrus"
)

type Register struct {
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	AuthProvider string `json:"authProvider" validate:"required,oneof=local google"`
	Password     string `json:"password" validate:"required_if=AuthProvider local,min=6"`
	GoogleID     string `json:"googleId" validate:"required_if=AuthProvider google"`
	GoogleToken  string `json:"googleToken" validate:"required_if=AuthProvider google"`
}

type RegisterHandler decorator.CommandHandler[Register]

type registerHandler struct {
	repo      userRepo.Repository
	validator validator.Validator
}

func NewRegisterHandler(
	userRepo userRepo.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
	validator validator.Validator,
) RegisterHandler {
	return decorator.ApplyCommandDecorators(
		registerHandler{repo: userRepo, validator: validator},
		logger,
		metricsClient,
	)
}

func (h registerHandler) Handle(ctx context.Context, cmd Register) error {
	fmt.Println("1")
	if err := h.validator.ValidateStruct(&cmd); err != nil {
		return err
	}

	fmt.Println("2")
	existingUser, err := h.repo.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return messages.ErrExists
	}
	fmt.Println("3")
	if existingUser != nil {
		return messages.ErrExists
	}
	fmt.Println("4")

	var hashedPassword string

	if cmd.AuthProvider == "local" {
		passwordHash, err := password.HashPassword(cmd.Password)
		if err != nil {
			return messages.ErrHashPassword
		}

		hashedPassword = passwordHash
	}

	user := &model.User{
		FirstName:    cmd.FirstName,
		LastName:     cmd.LastName,
		Email:        cmd.Email,
		Password:     hashedPassword,
		Status:       types.AccountStatus.PENDING,
		AuthProvider: cmd.AuthProvider,
		GoogleID:     cmd.GoogleID,
		GoogleToken:  cmd.GoogleToken,
	}

	return h.repo.Create(ctx, user)
}
