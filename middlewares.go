package main

import (
	"context"
	"fmt"

	"github.com/KrisQ/go-feed/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("you need to be logged int to follow a feed")
		}
		return handler(s, cmd, user)
	}
}
