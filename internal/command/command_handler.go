package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Tanay-Verma/gator/internal/database"
)

type commandHandler func(s *State, cmd Command) error

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("No arguments provided. A username is required.")
	}

	username := cmd.arguments[0]
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Username %q has been set.\n", username)

	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("No arguments provided. A username is required.")
	}

	username := cmd.arguments[0]
	newUser, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			Name:      username,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(newUser.Name)
	if err != nil {
		return err
	}

	fmt.Println("New User has been created. User:", newUser)
	return nil
}
