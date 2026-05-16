package main

import (
	"STA/handlers"
	"net/http"
	"time"
)

var client = http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/team/", handlers.GetTeam(client))

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

/*func team(w http.ResponseWriter, r *http.Request) {
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
*/
