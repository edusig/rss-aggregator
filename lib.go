package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode >= 500 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type returnError struct {
		Error string `json:"error"`
	}
	respondWithJson(w, statusCode, returnError{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(data)
}

func decodeJsonBody[T any](body io.ReadCloser, resp T) (T, error) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&resp)
	if err != nil {
		return resp, errors.New("couldn't decode parameters")
	}
	return resp, nil
}
