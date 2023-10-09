package main

import (
	"internal/database"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
}

func dbFeedToFeed(it database.Feed) Feed {
	return Feed{
		ID:        it.ID,
		Name:      it.Name,
		Url:       it.Url,
		CreatedAt: it.CreatedAt,
		UpdatedAt: it.UpdatedAt,
		UserID:    it.UserID,
	}
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func dbUserToUser(it database.User) User {
	return User{
		ID:        it.ID,
		CreatedAt: it.CreatedAt,
		UpdatedAt: it.UpdatedAt,
		Name:      it.Name,
		ApiKey:    it.ApiKey,
	}
}

type UsersFeed struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func dbUsersFeedToUsersFeed(it database.UsersFeed) UsersFeed {
	return UsersFeed{
		ID:        it.ID,
		CreatedAt: it.CreatedAt,
		UpdatedAt: it.UpdatedAt,
		UserID:    it.UserID,
		FeedID:    it.FeedID,
	}

}
