package main

import (
	"context"
	"fmt"
	"time"

	"github.com/KrisQ/go-feed/internal/database"
	"github.com/google/uuid"
)

func handleFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("this command doesn't have any arguments")
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't find the feed follows: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %v \n", feed.Feedname)
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("feed with url %v doesn't exist yet", url)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("couldn't create a new feed follow: %w", err)
	}

	fmt.Printf("* user name: %v \n", user.Name)
	fmt.Printf("* feed name: %v \n", feed.Name)

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Name: user.Name,
		Url:  url,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete a feed: %w", err)
	}
	return nil
}
