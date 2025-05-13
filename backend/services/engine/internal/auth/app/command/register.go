package command

import (
	"errors"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/types"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/decorator"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/model"
	userRepo "github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/repository"
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

func (h registerHandler) Handle(cmd Register) error {
	if err := h.validator.ValidateStruct(&cmd); err != nil {
		return err
	}

	existingUser, err := h.repo.GetByEmail(cmd.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	user := &model.User{
		FirstName:    cmd.FirstName,
		LastName:     cmd.LastName,
		Email:        cmd.Email,
		Password:     cmd.Password,
		Status:       types.AccountStatus.PENDING,
		AuthProvider: cmd.AuthProvider,
		GoogleID:     cmd.GoogleID,
		GoogleToken:  cmd.GoogleToken,
	}

	return h.repo.Create(user)
}
