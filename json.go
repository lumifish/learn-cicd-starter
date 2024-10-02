package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError) // Use http.StatusInternalServerError instead of 500
		return
	}
	w.WriteHeader(code)
	
	// Handle potential error from w.Write()
	_, err = w.Write(dat)
	if err != nil {
		// Log the error, but we can't really recover at this point
		log.Printf("Error writing response: %v", err)
		// Note: We've already sent headers, so we can't change the status code here
	}
}
