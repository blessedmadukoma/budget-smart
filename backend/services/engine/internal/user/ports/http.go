package ports

import (
	"log"
	"net/http"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/auth"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/middleware"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/app"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/user/app/query"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/json"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/messages"
	"github.com/go-chi/chi/v5"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(router *chi.Mux, app app.Application, a auth.AuthMiddleware, m middleware.HttpMiddleware) *HttpServer {
	s := &HttpServer{app: app}

	router.Route("/users", func(r chi.Router) {
		r.Use(a.IsAuthenticated, m.SingleRequest, m.Throttle(2))
		r.Get("/me", s.GetUser)
	})

	return s
}

func (h HttpServer) GetUser(w http.ResponseWriter, r *http.Request) {
	var c query.GetUser

	ctx := r.Context()

	accountKey, err := auth.GetAccount(ctx)
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, messages.ErrUnauthorized)
		return
	}

	c.UserID = accountKey.ID

	u, err := h.app.Queries.GetUser.Handle(ctx, c)
	if err != nil {
		log.Println("Error handling user query:", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, messages.OperationWasSuccessful, u)
}
