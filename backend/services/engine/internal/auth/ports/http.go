package ports

import (
	"context"
	"net/http"
	"time"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app/command"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app/query"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/json"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/messages"
	"github.com/go-chi/chi/v5"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(router *chi.Mux, app app.Application) *HttpServer {
	s := &HttpServer{app: app}

	router.Route("/auth", func(r chi.Router) {
		r.Post("/register", s.Register)
		r.Post("/login", s.Login)
	})

	return s
}

func (h HttpServer) Register(w http.ResponseWriter, r *http.Request) {
	var c command.Register

	if err := json.ParseJSON(r, &c); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	err := h.app.Commands.Register.Handle(ctx, c)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, messages.OperationWasSuccessful, nil)
}

func (h HttpServer) Login(w http.ResponseWriter, r *http.Request) {
	var c query.Login

	if err := json.ParseJSON(r, &c); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	token, err := h.app.Queries.Login.Handle(ctx, c)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	json.WriteJSON(w, http.StatusOK, messages.OperationWasSuccessful, response)
}
