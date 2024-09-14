package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiError struct {
	Message string `json:"message"`
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json response: %+v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(code)
	w.Write(res)
}

func respondWithErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Err Response with code %v, %v", code, msg)
	}

	respondWithJson(w, code, ApiError{Message: msg})
}
