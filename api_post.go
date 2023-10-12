package main

import (
	"internal/database"
	"log"
	"net/http"
)

func (cfg *apiConfig) getPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	dbPosts, err := cfg.DB.GetPostsByUser(r.Context(), user.ID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}
	respondWithJson(w, http.StatusOK, SliceMap(dbPosts, dbPostsToPosts))
}
