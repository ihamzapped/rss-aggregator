package main

import (
	"github.com/ihamzapped/rss-aggregator/internal/database"
	"time"
)

type ApiConfig struct {
	DB *database.Queries
}

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Language    string    `xml:"language"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

// RESOPONSE TYPES

type ErrResponse struct {
	Error string `json:"error"`
}

type PostResponse struct {
	database.FeedPost
	Description string `json:"description,omitempty"`
}

type FeedResponse struct {
	database.Feed
	LastFetchedAt *time.Time `json:"last_fetched_at,omitempty"`
}
