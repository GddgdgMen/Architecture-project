package main

import (
	"net/http"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var Time = time.Now()
}
