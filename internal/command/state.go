package command

import "github.com/Tanay-Verma/gator/internal/config"

type State struct {
	cfg *config.Config
}

func NewState(cfg *config.Config) State {
	return State{
		cfg: cfg,
	}
}
