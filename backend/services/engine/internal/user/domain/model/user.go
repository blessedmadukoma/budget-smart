package model

import (
	"time"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/auth/types"
)

type User struct {
	ID           uint                    `json:"-"`
	UID          string                  `json:"id,omitempty"`
	FirstName    string                  `json:"firstName"`
	LastName     string                  `json:"lastName"`
	Email        string                  `json:"email"`
	Password     string                  `json:"password"`
	Status       types.AccountStatusType `json:"status"`
	CreatedAt    time.Time               `json:"createdAt"`
	AuthProvider string                  `json:"authProvider"`
	GoogleID     string                  `json:"googleId"`
	GoogleToken  string                  `json:"googleToken"`
}

// Domain methods
func (u *User) IsActive() bool {
	return u.Status == "ACCEPTED"
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
