package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/KrisQ/go-feed/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <limit>", cmd.Name)
	}
	limit, err := strconv.Atoi(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid limit: %w", err)
	}

	if limit < 0 {
		return fmt.Errorf("limit must be a positive integer within the range of int32")
	}

	limitInt32 := int32(limit)
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: limitInt32,
	})
	for _, post := range posts {
		fmt.Print(post)
	}
	return nil
}
