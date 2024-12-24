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
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
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

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", u.Name)
			continue
		}
		fmt.Println("*", u.Name)
	}

	return nil
}

func HandlerReset(s *State, cmd Command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	return nil
}
