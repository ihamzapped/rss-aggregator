package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/ihamzapped/rss-aggregator/internal/auth"
	db "github.com/ihamzapped/rss-aggregator/internal/database"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, "All good!")
}

func errCheck(w http.ResponseWriter, r *http.Request) {
	respondErr(w, http.StatusBadRequest, "All Bad!")
}

func (s *ApiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	type body struct {
		Name string `json:"name"`
	}

	params, err := parseReq(w, r, body{})

	if err != nil {
		return
	}

	user, err := s.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:   uuid.New(),
		Name: params.Name,
	})

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Could not create usr: %v", err))
		return
	}

	respond(w, http.StatusCreated, user)
}

func (s *ApiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	key, err := auth.GetApiKey(r.Header)

	if err != nil {
		respondErr(w, http.StatusUnauthorized, fmt.Sprintf("%v", err))
		return
	}

	user, err := s.DB.GetUserByApiKey(r.Context(), key)

	if err != nil {
		respondErr(w, http.StatusNotFound, fmt.Sprintf("%v", err))
		return
	}

	respond(w, http.StatusOK, user)
}
