package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"STA/models"
)

const BASE_URL = "https://ftc-api.firstinspires.org/v2.0/2025"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /team/{teamNumber}", team)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	s.ListenAndServe()
}

func team(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	teamNumberStr := r.PathValue("teamNumber")

	teamNumber, err := strconv.Atoi(teamNumberStr)
	if err != nil {
		http.Error(w, "Invalid team number", http.StatusBadRequest)
		return
	}

	body, statusCode, err := apiRequest(
		fmt.Sprintf("teams?teamNumber=%d", teamNumber),
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if statusCode != http.StatusOK {
		http.Error(w, string(body), statusCode)
		return
	}

	var teamResponse models.TeamResponse

	if err := json.Unmarshal(body, &teamResponse); err != nil {
		http.Error(w, "Failed to decode team response", http.StatusBadGateway)
		return
	}

	if len(teamResponse.Teams) == 0 {
		http.Error(w, "Team not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(teamResponse.Teams[0]); err != nil {
		http.Error(w, "Failed to encode team response", http.StatusInternalServerError)
		return
	}
}

func apiRequest(path string) ([]byte, int, error) {
	client := http.Client{
		Timeout: time.Minute * 2,
	}

	targetURL := fmt.Sprintf("%s/%s", BASE_URL, path)

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("User-Agent", "STA")

	req.SetBasicAuth(
		os.Getenv("FTC_EVENTS_USERNAME"),
		os.Getenv("FTC_EVENTS_AUTH_TOKEN"),
	)

	res, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, res.StatusCode, nil
}
