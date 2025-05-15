package auth

import (
	"database/sql"

	"github.com/blessedmadukoma/budgetsmart/engine/config"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app/command"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app/query"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/ports"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/cache"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/metrics"
	userRepo "github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/repository"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/validator"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func NewService(router *chi.Mux, db *sql.DB, cache cache.Cache) error {

	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	userRepo := userRepo.NewRepository(db)

	validator := validator.NewValidator()

	ports.NewHttpServer(router, app.Application{
		Commands: app.Commands{
			Register: command.NewRegisterHandler(userRepo, logger, metricsClient, validator),
		},
		Queries: app.Queries{
			Login: query.NewLoginHandler(userRepo, logger, metricsClient, validator, config.Envs, cache),
		},
	})

	return nil
}
