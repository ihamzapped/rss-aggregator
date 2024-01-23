package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ihamzapped/rss-aggregator/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenReq time.Duration) {
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenReq)

	ticker := time.NewTicker(timeBetweenReq)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurrency))

		if err != nil {
			log.Println("Error fetching feeds:", err)
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapper(feed, wg, db)
		}
	}
}

func scrapper(feed database.Feed, wg *sync.WaitGroup, db *database.Queries) {
	defer wg.Done()

	_, err := db.UpdateLastFetch(context.Background(), feed.ID)

	if err != nil {
		log.Println("Error UpdateLastFetch:", err)
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		desc := sql.NullString{}

		if item.Description != "" {
			desc.String = item.Description
			desc.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			log.Printf("Error parsing date %v, error: %v", item.PubDate, err)
		}

		err = db.CreateFeedPost(context.Background(), database.CreateFeedPostParams{
			ID:          uuid.New(),
			FeedID:      feed.ID,
			Title:       item.Title,
			Description: desc,
			PublishedAt: pubAt,
			Url:         item.Link,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Failed to create post err: %v", err)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}
