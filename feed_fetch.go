package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"internal/database"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

func FeedWorker(db *database.Queries) {
	log.Println("Starting Feed Worker")
	for {
		log.Println("Feed Worker - Fetching the next 10 feeds")
		rows, err := db.GetNextFeedsToFetch(context.TODO())
		if err != nil {
			break
		}

		var wg sync.WaitGroup
		for _, row := range rows {
			wg.Add(1)
			go func(row database.GetNextFeedsToFetchRow) {
				defer wg.Done()
				err := fetchFeed(row.Url, row.ID, db)
				if err != nil {
					log.Fatal(err)
					log.Fatalf("Error while fetching feed from %v", row.Url)
				}
			}(row)
		}
		wg.Wait()

		time.Sleep(30 * time.Minute)
	}
}

func fetchFeed(url string, id uuid.UUID, db *database.Queries) error {
	rss, err := fetchRSS(url)
	if err != nil {
		return err
	}
	log.Printf("Fetched %v from %v", rss.Channel.Title, url)
	for _, item := range rss.Channel.Items {
		log.Printf("Found item %v from %v", *item.Title, rss.Channel.Title)
	}
	db.MarkFeedFetched(context.TODO(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            id,
	})
	return nil
}

func fetchRSS(url string) (RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return RSS{}, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
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
