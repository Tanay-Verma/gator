package command

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Tanay-Verma/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.arguments) == 1 {
		if sepcifiedLimit, err := strconv.Atoi(cmd.arguments[0]); err == nil {
			limit = sepcifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w\n", err)
		}
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w\n", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("  %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("========================================")
	}

	return nil
}
