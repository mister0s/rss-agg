package main

import "net/http"

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
