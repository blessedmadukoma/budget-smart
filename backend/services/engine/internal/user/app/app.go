package app

import "github.com/blessedmadukoma/budgetsmart/engine/internal/user/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	GetUser query.GetUserHandler
}
