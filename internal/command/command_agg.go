package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Tanay-Verma/gator/internal/database"
	"github.com/Tanay-Verma/gator/internal/rss"
	"github.com/lib/pq"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Usage: agg <time_between_reqs>")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("Collecting feeds every", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *State) {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s, feed)
}

func scrapeFeed(s *State, feed database.GetNextFeedToFetchRow) {
	err := s.Db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now().UTC(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	})
	if err != nil {
		log.Printf("Couldn't mark feed %q fetched: %v\n", feed.Name, err)
		return
	}

	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %q: %v\n", feed.Name, err)
		return
	}

	fmt.Println("-", rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		publisedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publisedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := s.Db.CreatePost(context.Background(), database.CreatePostParams{
			Title: item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publisedAt,
			FeedID:      feed.ID,
		})
		if err != nil {

			// Error Code: "23505": "unique_violation"
			if err, ok := err.(*pq.Error); ok && err.Code == "23505" {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", rssFeed.Channel.Title, len(rssFeed.Channel.Item))
}
