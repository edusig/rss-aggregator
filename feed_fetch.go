package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"internal/database"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

var ErrMissingPostTitle = errors.New("missing post title")
var ErrMissingPostUrl = errors.New("missing post url")

func FeedWorker(db *database.Queries) {
	log.Println("Starting Feed Worker")
	for {
		log.Println("Feed Worker - Fetching the next 10 feeds")
		rows, err := db.GetNextFeedsToFetch(context.TODO())
		if err != nil {
			break
		}
		log.Println(rows)

		var wg sync.WaitGroup
		for _, row := range rows {
			wg.Add(1)
			go func(row database.GetNextFeedsToFetchRow) {
				defer wg.Done()
				err := fetchFeed(row.Url, row.ID, db)
				if err != nil {
					log.Fatalf("Error while fetching feed from %v \n%v", row.Url, err)
				}
			}(row)
		}
		wg.Wait()

		time.Sleep(1 * time.Minute)
	}
}

func fetchFeed(url string, id uuid.UUID, db *database.Queries) error {
	rss, err := fetchRSS(url)
	if err != nil {
		return err
	}
	for _, item := range rss.Channel.Items {
		postCreate, err := feedItemToPostCreate(id, item)
		if err != nil {
			continue
		}
		db.CreatePost(context.TODO(), postCreate)
	}
	db.MarkFeedFetched(context.TODO(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            id,
	})
	log.Printf("Fetched %v from %v", rss.Channel.Title, url)
	return nil
}

func feedItemToPostCreate(feedId uuid.UUID, item RSSItem) (database.CreatePostParams, error) {
	if item.Title == nil {
		return database.CreatePostParams{}, ErrMissingPostTitle
	}
	if item.Link == nil {
		return database.CreatePostParams{}, ErrMissingPostUrl
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return database.CreatePostParams{}, err
	}

	description := sql.NullString{String: "", Valid: false}
	if item.Description != nil {
		description = sql.NullString{String: *item.Description, Valid: true}
	}
	pubDate := sql.NullTime{Time: time.Time{}, Valid: false}
	if item.PubDate != nil {
		parsed, err := time.Parse(time.RFC822, *item.PubDate)
		if err == nil {
			pubDate = sql.NullTime{Time: parsed, Valid: true}
		}
	}

	return database.CreatePostParams{
		ID:          newUUID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       *item.Title,
		Url:         *item.Link,
		Description: description,
		PublishedAt: pubDate,
		FeedID:      feedId,
	}, nil
}

func fetchRSS(url string) (RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return RSS{}, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n%v\n", resp.StatusCode, body, url)
	}
	if err != nil {
		return RSS{}, err
	}
	rss := RSS{}
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return rss, err
	}
	return rss, nil
}

type RSSResponse struct {
	RSS RSS `xml:"rss"`
}

type RSS struct {
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Items       []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title       *string `xml:"title"`
	Link        *string `xml:"link"`
	Description *string `xml:"description"`
	PubDate     *string `xml:"pubDate"`
}
