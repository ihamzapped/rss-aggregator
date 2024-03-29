package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
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

	params, err := parseBody(w, r, body{})

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

func (api *ApiConfig) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := api.DB.GetAllUsers(context.Background())

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Could not fetch users: %v", err))
		return
	}

	respond(w, http.StatusOK, users)
}

func (api *ApiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type body struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params, err := parseBody(w, r, body{})

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

	respond(w, http.StatusOK, createFeedsResponse(feeds))
}

func (api *ApiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type body struct {
		Feed_id uuid.UUID `json:"feed_id"`
	}

	params, err := parseBody(w, r, body{})

	if err != nil {
		return
	}

	feed, err := api.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		FeedID: params.Feed_id,
		UserID: user.ID,
	})

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Could not create follow: %v", err))
		return
	}

	respond(w, http.StatusCreated, feed)
}

func (api *ApiConfig) handleGetUserFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	follows, err := api.DB.GetUserFollows(r.Context(), user.ID)

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Could not fetch: %v", err))
		return
	}

	respond(w, http.StatusOK, follows)
}

func (api *ApiConfig) handleDeleteUserFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	followIDStr := chi.URLParam(r, "follow_id")
	followID, err := uuid.Parse(followIDStr)

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Could not delete follow: %v", err))
		return
	}

	err = api.DB.DeleteFollow(r.Context(), database.DeleteFollowParams{
		ID:     followID,
		UserID: user.ID,
	})

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Could not delete follow: %v", err))
		return
	}

	respond(w, http.StatusOK, struct{}{})
}

func (api *ApiConfig) handleGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	posts, err := api.DB.GetPostsForUser(r.Context(), user.ID)

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Could not fetch posts for user: %v", err))
		return
	}

	respond(w, http.StatusOK, createPostsResponse(posts))
}
