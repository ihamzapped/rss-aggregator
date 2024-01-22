package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ihamzapped/rss-aggregator/internal/database"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, "All good!")
}

func errCheck(w http.ResponseWriter, r *http.Request) {
	respondErr(w, http.StatusBadRequest, "All Bad!")
}

func (api *ApiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	type body struct {
		Name string `json:"name"`
	}

	params, err := parseReq(w, r, body{})

	if err != nil {
		return
	}

	user, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:   uuid.New(),
		Name: params.Name,
	})

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Could not create usr: %v", err))
		return
	}

	respond(w, http.StatusCreated, user)
}

func (api *ApiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respond(w, http.StatusOK, user)
}

func (api *ApiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type body struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params, err := parseReq(w, r, body{})

	if err != nil {
		return
	}

	feed, err := api.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:     uuid.New(),
		Url:    params.Url,
		Name:   params.Name,
		UserID: user.ID,
	})

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Could not create feed: %v", err))
		return
	}

	respond(w, http.StatusCreated, feed)
}

func (api *ApiConfig) handleGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := api.DB.GetAllFeeds(r.Context())

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	respond(w, http.StatusOK, feeds)
}
