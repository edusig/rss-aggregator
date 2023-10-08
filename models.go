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

func dbFeedToFeed(f database.Feed) Feed {
	return Feed{
		ID:        f.ID,
		Name:      f.Name,
		Url:       f.Url,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		UserID:    f.UserID,
	}
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func dbUserToUser(u database.User) User {
	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		ApiKey:    u.ApiKey,
	}
}
