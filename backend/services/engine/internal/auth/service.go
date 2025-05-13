package auth

import (
	"database/sql"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app/command"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/ports"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/metrics"
	userRepo "github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/repository"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/validator"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func NewService(router *chi.Mux, db *sql.DB) error {

	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	userRepo := userRepo.NewRepository(db)

	validator := validator.NewValidator()

	ports.NewHttpServer(router, app.Application{
		Commands: app.Commands{
			Register: command.NewRegisterHandler(userRepo, logger, metricsClient, validator),
		},
	})

	return nil
}
