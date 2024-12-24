package command

import (
	"github.com/Tanay-Verma/gator/internal/config"
	"github.com/Tanay-Verma/gator/internal/database"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
}

func NewState(cfg *config.Config, db *database.Queries) State {
	return State{
		cfg: cfg,
		db:  db,
	}
}
