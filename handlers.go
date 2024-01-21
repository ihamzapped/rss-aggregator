package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	db "github.com/ihamzapped/rss-aggregator/internal/database"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, "All good!")
}

func errCheck(w http.ResponseWriter, r *http.Request) {
	respondErr(w, http.StatusBadRequest, "All Bad!")
}

func (s *ApiConfig) createUser(w http.ResponseWriter, r *http.Request) {

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

	respond(w, http.StatusOK, user)
}
