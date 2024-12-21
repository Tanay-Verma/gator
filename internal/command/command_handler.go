package command

import (
	"errors"
	"fmt"
)

type commandHandler func(s *State, cmd Command) error

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("No arguments provided. A username is required.")
	}

	username := cmd.arguments[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Username %q has been set.\n", username)

	return nil
}
