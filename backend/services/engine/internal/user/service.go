package user

import (
	"database/sql"

	"github.com/blessedmadukoma/budgetsmart/engine/config"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/auth"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/cache"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/middleware"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/app"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/app/query"
	userRepo "github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/repository"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/ports"
	"github.com/go-chi/chi/v5"
)

func NewService(router *chi.Mux, db *sql.DB, cache cache.Cache, config config.Config) error {

	// logger := logrus.NewEntry(logrus.StandardLogger())
	// metricsClient := metrics.NoOp{}

	userRepo := userRepo.NewRepository(db)

	ports.NewHttpServer(router, app.Application{
		Queries: app.Queries{
			GetUser: query.NewGetUserHandler(userRepo),
		},
	}, auth.NewAuthMiddleware(userRepo, cache), middleware.NewHttpMiddleware(config))

	return nil
}
