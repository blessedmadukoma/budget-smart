package ports

import (
	"net/http"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app/command"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/json"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/messages"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(router *chi.Mux, app app.Application) *HttpServer {
	s := &HttpServer{app: app}

	router.Route("/auth", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Post("/register", s.Register)
	})

	return s
}

func (h HttpServer) Register(w http.ResponseWriter, r *http.Request) {
	var c command.Register

	if err := json.ParseJSON(r, &c); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.app.Commands.Register.Handle(c)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, messages.ErrBadRequest)
		return
	}

	json.WriteJSON(w, http.StatusCreated, messages.OperationWasSuccessful, nil)
}
