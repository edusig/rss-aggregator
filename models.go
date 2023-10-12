package main

import (
	"internal/database"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID  `json:"uuid"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func dbFeedToFeed(it database.Feed) Feed {
	return Feed{
		ID:            it.ID,
		Name:          it.Name,
		Url:           it.Url,
		CreatedAt:     it.CreatedAt,
		UpdatedAt:     it.UpdatedAt,
		UserID:        it.UserID,
		LastFetchedAt: &it.LastFetchedAt.Time,
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

type Post struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description *string
	PublishedAt *time.Time
	FeedID      uuid.UUID
}

func dbPostsToPosts(it database.Post) Post {
	return Post{
		ID:          it.ID,
		CreatedAt:   it.CreatedAt,
		UpdatedAt:   it.UpdatedAt,
		Title:       it.Title,
		Url:         it.Url,
		Description: &it.Description.String,
		PublishedAt: &it.PublishedAt.Time,
		FeedID:      it.FeedID,
	}
}
