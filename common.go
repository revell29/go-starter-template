package main

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	JSONResponseWriter(w, http.StatusOK, map[string]string{"status": "ok"})
}

func JSONResponseWriter(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
