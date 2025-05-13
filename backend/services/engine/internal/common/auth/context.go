package auth

import (
	"context"
	"errors"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/types"
)

type Account struct {
	ID     uint
	Status types.AccountStatusType
}

type ctxKey int

const (
	accountContextKey ctxKey = iota
)

func GetAccount(ctx context.Context) (Account, error) {
	a, ok := ctx.Value(accountContextKey).(Account)
	if ok {
		return a, nil
	}

	return Account{}, errors.New("no user in context")
}

func setAccount(ctx context.Context, a Account) context.Context {
	return context.WithValue(ctx, accountContextKey, a)
}
