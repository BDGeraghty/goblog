package main

import (
	"context"
	"fmt"
	"time"
	"log"
)

func handlerAgg(s *state, cmd command) error {
	/*
	err := scrapeFeeds(s)
	if err != nil {
		return fmt.Errorf("couldn't scrape feeds: %w", err)
	}
	*/
		
	/*
	feed, err := fetchFeed(context.Background(), "https://blog.boot.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	*/
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	// Infinite loop - this function will never return unless there's an error earlier
	for range ticker.C {
		if err := scrapeFeeds(s); err != nil {
			log.Printf("Error scraping feeds: %v", err)
		}
	}
	return nil
	
}

func scrapeFeeds(s *state) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		_, err := s.db.GetNextFeedToFetch(context.Background())
		if err != nil {
			fmt.Printf("Error getting next feed to fetch: %v\n", err)
			continue
		}
		/*
		err = scrapeFeeds(s, cmd)
		if err != nil {
			fmt.Printf("Error scraping feed %s: %v\n", feed.Name, err)
			continue
		}
		*/
		// Mark the feed as fetched
		err = s.db.MarkFeedFetched(context.Background(), feed.ID)
		if err != nil {
			fmt.Printf("Error marking feed %s as fetched: %v\n", feed.Name, err)
		}
		fmt.Printf("Feed %s scraped successfully.\n", feed.Name)
		// Here you would typically process the feed items, e.g., save them to the database
		items, err := fetchFeedPosts(context.Background(), feed.Url)
		if err != nil {
			fmt.Printf("Error fetching items for feed %s: %v\n", feed.Name, err)
			continue
		}	
		for _, item := range items {
			// Here you would typically save the item to the database
			fmt.Printf("Item: %s - %s\n", item.Title, item.Link)
			// You can also print the item description if needed
			// fmt.Printf("Description: %s\n", item.Description)	
			// fmt.Printf("Published: %s\n", item.PubDate)
			
		}	
	}

	return nil
}

