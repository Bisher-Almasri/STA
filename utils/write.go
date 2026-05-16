package utils

import (
	"STA/models"
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
	}
}

func WriteError(w http.ResponseWriter, message string, status int) {
	WriteJSON(w, status, models.ErrorResponse{
		Error:  message,
		Status: status,
	})
}
