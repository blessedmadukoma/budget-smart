package api

import (
	"database/sql"
	"net/http"

	"github.com/blessedmadukoma/budgetsmart/engine/config"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/cache"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/middleware"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/log"
	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	addr   string
	config config.Config
	db     *sql.DB
	cache  cache.Cache
	logger *log.Logger
}

func NewAPIServer(addr string, config config.Config, db *sql.DB, cache cache.Cache, logger *log.Logger) *APIServer {
	return &APIServer{
		addr:   addr,
		config: config,
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

func (s *APIServer) setupMiddleware() *chi.Mux {

	router := chi.NewRouter()

	m := middleware.NewHttpMiddleware(s.config)

	router.Use(func(next http.Handler) http.Handler {
		return s.logger.Middleware(next)
	})

	router.Use(func(next http.Handler) http.Handler {
		return m.CORS(next)
	})

	router.Use(func(next http.Handler) http.Handler {
		return m.EnhanceContext(next)
	})

	return router
}

func (s *APIServer) Run() error {
	s.logger.SetPrefix("api")
	router := s.setupMiddleware()

	// add sentry or datadog or APIToolKit

	subrouter := chi.NewRouter()
	router.Mount("/api/v1", subrouter)

	auth.NewService(subrouter, s.db, s.cache)
	user.NewService(subrouter, s.db, s.cache, s.config)

	s.logger.Printf("Listening on: %s", s.addr)

	return http.ListenAndServe(s.addr, router)
}
