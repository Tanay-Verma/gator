package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Tanay-Verma/gator/internal/database"
	"github.com/Tanay-Verma/gator/internal/rss"
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

func HandlerAgg(s *State, cmd Command) error {
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Println(rssFeed)
	return nil
}

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return errors.New("Usage: addfeed <feed_name> <feed_url>")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %v\n", feed)

	return nil
}

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("* Feed Name: %s | Feed URL: %s | User Name: %s\n", feed.FeedName, feed.Url, feed.UserName)
	}
	return nil
}

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Usage: follow <feed_url>")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	new_feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed Name: %s | Follower Name: %s",
		new_feed_follow.FeedName,
		new_feed_follow.UserName,
	)

	return nil
}

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Feeds Followed by %q:\n", s.cfg.CurrentUserName)
	for _, f := range feedsFollowed {
		fmt.Printf("* %s\n", f.FeedName)
	}
	return nil
}

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Usage: unfollow <feed_url>")
	}

	err := s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		Url:    cmd.arguments[0],
	})
	if err != nil {
		return err
	}

	return nil
}
