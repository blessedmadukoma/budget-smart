package app

import "github.com/blessedmadukoma/budgetsmart/engine/internal/auth/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Register command.RegisterHandler
}

type Queries struct {
}
