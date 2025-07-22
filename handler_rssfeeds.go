package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

// RSSFeed represents the structure of an RSS feed
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

// RSSItem represents an individual item in an RSS feed
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}	
	if feedURL == "" {
		return nil, fmt.Errorf("invalid feed URL")
	}	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't create request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read response body: %w", err)
	}

	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshal RSS feed: %w", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Items  {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Items[i] = item
	}

	return &feed, nil
}
/*
func handlerRSSFeeds(s *state, cmd command) error {
	
	cmd.Args = append(cmd.Args, "https://www.wagslane.dev/index.xml") // Temporary Hard Coded URL
	
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed-url>", cmd.Name)
	}

	feedURL := cmd.Args[0]
	
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't fetch RSS feed: %w", err)
	}

	fmt.Printf("RSS Feed Fetched Successfully from %s:\n", feedURL)
	printRSSFeed(feed)
	return nil
}
*/
/*
func printRSSFeed(feed *RSSFeed) {
	
	
	fmt.Printf(" * Title:       %v\n", feed.Channel.Title)
	fmt.Printf(" * Link:        %v\n", feed.Channel.Link)
	fmt.Printf(" * Description: %v\n", feed.Channel.Description)
	fmt.Printf(" * Items:       %d\n", len(feed.Channel.Items))

	

	//feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	//feed.Channel.Link = html.UnescapeString(feed.Channel.Link)
	//feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	

	for _, item := range feed.Channel.Items {
		item.Title = html.UnescapeString(item.Title)
		//item.Link = html.UnescapeString(item.Link)
		item.Description = html.UnescapeString(item.Description)
		//item.PubDate = html.UnescapeString(item.PubDate)
		fmt.Printf("   - %s (%s)\n", item.Title, item.Link)
	}
	//fmt.Printf("   - %s (%s)\n", "Optimize for simplicity", "https://wagslane.dev/posts/optimize-for-simplicit-first/)")
	
	fmt.Printf("%+v\n", feed)
}
*/	
