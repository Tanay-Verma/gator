package command

import (
	"context"

	"github.com/Tanay-Verma/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) commandHandler {
	return func(s *State, cmd Command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		err = handler(s, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
}
