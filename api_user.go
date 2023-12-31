package main

import (
	"internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) userCreateHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params, err := decodeJsonBody(r.Body, parameters{})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	userId, err := uuid.NewRandom()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate an user id")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        userId,
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
	}
	respondWithJson(w, http.StatusCreated, dbUserToUser(user))
}

func (cfg *apiConfig) userGetHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, http.StatusCreated, dbUserToUser(user))
}
