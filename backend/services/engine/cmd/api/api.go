package api

import (
	"database/sql"
	"net/http"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/middleware"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/log"
	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	addr   string
	db     *sql.DB
	logger *log.Logger
}

func NewAPIServer(addr string, db *sql.DB, logger *log.Logger) *APIServer {
	return &APIServer{
		addr:   addr,
		db:     db,
		logger: logger,
	}
}

func (s *APIServer) setupMiddleware() *chi.Mux {

	router := chi.NewRouter()

	m := middleware.NewHttpMiddleware()

	router.Use(func(next http.Handler) http.Handler {
		return m.CORS(next)
	})

	return router
}

func (s *APIServer) Run() error {
	s.logger.SetPrefix("api")
	router := s.setupMiddleware()

	// add sentry or datadog or APIToolKit

	subrouter := chi.NewRouter()
	router.Mount("/api/v1", subrouter)

	auth.NewService(subrouter, s.db)

	s.logger.Printf("Listening on: %s", s.addr)

	return http.ListenAndServe(s.addr, router)
}
