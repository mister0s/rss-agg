package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mister0s/rss-agg/internal/auth"
	"github.com/mister0s/rss-agg/internal/database"
)

type createUserReq struct {
	Name string `json:"name"`
}

func (s *ApiServer) createUserHandler(w http.ResponseWriter, r *http.Request) {
	req := createUserReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Failed to parse create user params: %v", err))
		return
	}
	if req.Name == "" {
		respondWithErr(w, 400, "Please input user name")
		return
	}

	newUser, err := s.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      req.Name,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Failed to create user: %v", err))
	}

	respondWithJson(w, 201, dbUserToUser(newUser))
}

func (s *ApiServer) getUserHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)

	if err != nil {
		respondWithErr(w, 403, fmt.Sprintf("Auth Failed: %v", err))
		return
	}

	if apiKey == "" {
		respondWithErr(w, 403, "Auth Failed: Empty api key")
		return
	}

	user, err := s.DB.GetUserByApiKey(r.Context(), apiKey)

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Failed to get user: %v", err))
		return
	}

	respondWithJson(w, 201, dbUserToUser(user))
}
