package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Tanay-Verma/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Usage: follow <feed_url>")
	}

	feed, err := s.Db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	new_feed_follow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
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
	feedsFollowed, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Feeds Followed by %q:\n", s.Cfg.CurrentUserName)
	for _, f := range feedsFollowed {
		fmt.Printf("* %s\n", f.FeedName)
	}
	return nil
}

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Usage: unfollow <feed_url>")
	}

	err := s.Db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		Url:    cmd.arguments[0],
	})
	if err != nil {
		return err
	}

	return nil
}
