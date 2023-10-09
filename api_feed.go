package main

import (
	"internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	type response struct {
		Feed       Feed      `json:"feed"`
		FeedFollow UsersFeed `json:"feed_follow"`
	}

	params, err := decodeJsonBody(r.Body, parameters{})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate an feed id")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        newUUID,
		Name:      params.Name,
		Url:       params.Url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
	}

	newUUID, err = uuid.NewRandom()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate an feedFollow id")
		return
	}

	feedFollow, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        newUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't follow created feed")
	}

	respondWithJson(w, http.StatusCreated, response{
		Feed:       dbFeedToFeed(feed),
		FeedFollow: dbUsersFeedToUsersFeed(feedFollow),
	})
}

func (cfg *apiConfig) getAllFeedsHandler(w http.ResponseWriter, r *http.Request) {
	dbFeeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
	}
	respondWithJson(w, http.StatusOK, SliceMap(dbFeeds, dbFeedToFeed))
}
