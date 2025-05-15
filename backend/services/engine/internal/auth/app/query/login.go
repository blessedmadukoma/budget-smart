package query

import (
	"context"
	"fmt"
	"time"

	"github.com/blessedmadukoma/budgetsmart/engine/config"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/cache"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/decorator"
	userRepo "github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/repository"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/jwt"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/messages"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/password"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/validator"
	"github.com/sirupsen/logrus"
)

type Login struct {
	Email        string `json:"email" validate:"required,email"`
	AuthProvider string `json:"authProvider" validate:"required,oneof=local google"`
	Password     string `json:"password" validate:"required_if=AuthProvider local,min=6"`
	GoogleID     string `json:"googleId" validate:"required_if=AuthProvider google"`
	GoogleToken  string `json:"googleToken" validate:"required_if=AuthProvider google"`
}

type LoginHandler decorator.QueryHandler[Login, string]

type loginHandler struct {
	repo      userRepo.Repository
	validator validator.Validator
	config    config.Config
	cache     cache.Cache
}

func NewLoginHandler(
	userRepo userRepo.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
	validator validator.Validator,
	cfg config.Config,
	cache cache.Cache,
) LoginHandler {
	return decorator.ApplyQueryDecorators(
		&loginHandler{repo: userRepo, validator: validator, config: cfg, cache: cache},
		logger,
		metricsClient,
	)
}

func (h loginHandler) Handle(ctx context.Context, cmd Login) (string, error) {
	if err := h.validator.ValidateStruct(&cmd); err != nil {
		return "", err
	}

	user, err := h.repo.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return "", fmt.Errorf("user account %s", messages.ErrNotExists)
	}
	if user == nil {
		return "", fmt.Errorf("user account %s", messages.ErrNotExists)
	}

	switch cmd.AuthProvider {
	case "local":
		if !password.ComparePasswords(user.Password, []byte(cmd.Password)) {
			fmt.Println("compare password issues...")
			return "", messages.ErrInvalidPassword
		}
	case "google":
		if user.GoogleID != cmd.GoogleID {
			return "", messages.ErrInvalidCredentials
		}
		// Optionally, verify GoogleToken here if needed
	default:
		return "", fmt.Errorf("unsupported auth provider")
	}

	clientIP := jwt.GetToken(ctx)
	token, err := jwt.Create(user.ID, clientIP, h.config.JWTSecret, int(h.config.JWTExpirationInSeconds))
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	// store the token in cache
	h.cache.Set(context.Background(), cache.JWTTokenKey(user.UID), &token, time.Hour*1440)

	return token, nil
}
