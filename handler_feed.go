package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bdgeraghty/GoBlog/internal/database"
	"github.com/google/uuid"
	//"golang.org/x/tools/go/cfg"
)

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}
	
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	
	name := cmd.Args[0]
	url := cmd.Args[1]

		feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	
	fmt.Println("Feeds:")
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user for feed %s: %w", feed.Name, err)
		}
		printFeed(feed, user)
		fmt.Println()
	}
	fmt.Println(feeds)

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf(" * ID:          %v\n", feed.ID)
	fmt.Printf(" * Name:        %v\n", feed.Name)
	fmt.Printf(" * URL:         %v\n", feed.Url)
	fmt.Printf(" * User ID:     %v\n", feed.UserID)
	fmt.Printf(" * Created At:  %v\n", feed.CreatedAt)
	fmt.Printf(" * Updated At:  %v\n", feed.UpdatedAt)
	fmt.Printf(" * User Name:   %v\n", user.Name)
}