package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/types"
	"github.com/blessedmadukoma/budgetsmart/engine/internal/common/cache"
	userRepo "github.com/blessedmadukoma/budgetsmart/engine/internal/user/domain/repository"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/json"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/jwt"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/messages"
)

type AuthMiddleware struct {
	userRepo userRepo.Repository
	cache    cache.Cache
}

func NewAuthMiddleware(userRepo userRepo.Repository, cache cache.Cache) AuthMiddleware {
	return AuthMiddleware{userRepo: userRepo, cache: cache}
}

func (m AuthMiddleware) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.New(r)
		if err != nil {
			json.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		a, err := m.userRepo.GetByID(token.ID)
		if err != nil {
			json.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		var existingToken string

		err = m.cache.Get(context.Background(), cache.JWTTokenKey(a.UID), &existingToken)
		if err != nil {
			json.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		if existingToken != token.Raw {
			json.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		// Set account in context
		ctx := setAccount(r.Context(), Account{
			ID:     token.ID,
			Status: a.Status,
		})

		// Continue with the request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m AuthMiddleware) IsActive(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a, err := GetAccount(r.Context())
		if err != nil {
			json.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		if a.Status != types.AccountStatus.ACTIVE && a.Status != types.AccountStatus.RESTRICTED {
			json.WriteError(w, http.StatusUnauthorized, messages.WrapError(nil, messages.AccountPending))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m AuthMiddleware) IsLocked(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a, err := GetAccount(r.Context())
		if err != nil {
			json.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		if a.Status == types.AccountStatus.LOCKED {
			json.WriteError(w, http.StatusUnauthorized, messages.WrapError(nil, messages.AccountIsLocked))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m AuthMiddleware) IsRestricted(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a, err := GetAccount(r.Context())
		if err != nil {
			json.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		if a.Status == types.AccountStatus.RESTRICTED {
			json.WriteError(w, http.StatusUnauthorized, messages.WrapError(nil, messages.AccountIsRestricted))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m AuthMiddleware) IsAdminAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-KEY")

		if apiKey != os.Getenv("ADMIN_API_KEY") {
			json.WriteError(w, http.StatusUnauthorized, messages.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
