package main

import (
	"github.com/ihamzapped/rss-aggregator/internal/database"
)

type ErrResponse struct {
	Error string `json:"error"`
}

type ApiConfig struct {
	DB *database.Queries
}
