package main

import (
	"context"
	"internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) userCreateHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	type response struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
	}

	params := parameters{}
	params, err := decodeJsonBody(r.Body, params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	userId, err := uuid.NewRandom()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate an user id")
		return
	}

	user, err := cfg.DB.CreateUser(context.TODO(), database.CreateUserParams{
		ID:        userId,
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
	}
	respondWithJson(w, http.StatusCreated, response{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
	})
}
