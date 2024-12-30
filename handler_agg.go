package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/KrisQ/go-feed/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	currentTime := time.Now()

	nullTime := sql.NullTime{
		Time:  currentTime,
		Valid: true,
	}

	err := db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: nullTime,
	})
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Saving post: %s\n", item.Title)

		nullDescription := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}

		var parsedTime time.Time
		var validTime bool
		if item.PubDate != "" {
			parsedTime, err = time.Parse(time.RFC1123, item.PubDate)
			if err == nil {
				validTime = true
			} else {
				log.Printf("Failed to parse pubdate '%s': %v", item.PubDate, err)
			}
		}

		nullPubDate := sql.NullTime{
			Time:  parsedTime,
			Valid: validTime,
		}
		_, err := db.AddPost(context.Background(), database.AddPostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: nullDescription,
			PublishedAt: nullPubDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			log.Printf("Couldn't save post %s: %v", item.Title, err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
