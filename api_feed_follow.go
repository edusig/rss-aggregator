package main

import (
	"internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) followFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId string `json:"feed_id"`
	}

	params, err := decodeJsonBody(r.Body, parameters{})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedId, err := uuid.Parse(params.FeedId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed id")
		return
	}

	feed, err := cfg.DB.GetFeed(r.Context(), feedId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Feed not found")
		return
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate an feed id")
		return
	}

	follow, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        newUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't follow the feed")
		return
	}

	respondWithJson(w, http.StatusCreated, dbUsersFeedToUsersFeed(follow))
}

func (cfg *apiConfig) unfollowFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId string `json:"feed_id"`
	}

	params, err := decodeJsonBody(r.Body, parameters{})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedId, err := uuid.Parse(params.FeedId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed id")
		return
	}

	feed, err := cfg.DB.GetFeed(r.Context(), feedId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Feed not found")
		return
	}

	err = cfg.DB.UnfollowFeed(r.Context(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't unfollow the feed")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (cfg *apiConfig) getFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	follows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get user the feed follows")
		return
	}
	respondWithJson(w, http.StatusOK, SliceMap(follows, dbUsersFeedToUsersFeed))
}
