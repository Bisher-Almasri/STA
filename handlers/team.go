package handlers

import (
	"STA/models"
	"STA/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Get team
// @Description Returns a team by number
// @Tags Teams
// @Produce json
// @Param number path int true "Team Number"
// @Success 200 {object} models.TeamResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/teams/{number} [get]
func GetTeam(client http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const prefix = "/api/v1/teams/"
		if !strings.HasPrefix(r.URL.Path, prefix) {
			http.NotFound(w, r)
			return
		}

		teamStr := strings.TrimPrefix(r.URL.Path, prefix)
		teamNumber, err := strconv.Atoi(teamStr)
		if err != nil {
			utils.WriteError(w, "invalid team number", http.StatusBadRequest)
			return
		}

		url := fmt.Sprintf("teams?teamNumber=%d", teamNumber)

		bodyBytes, err := utils.ApiRequest(url, client)
		if err != nil {
			log.Printf("ApiRequest error: %v", err)
			utils.WriteError(w, "upstream API error", http.StatusBadGateway)
			return
		}

		var response models.TeamResponse
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			log.Printf("decode error: %v", err)
			utils.WriteError(w, "failed to decode team response", http.StatusBadGateway)
			return
		}

		if len(response.Teams) == 0 {
			utils.WriteError(w, "team not found", http.StatusNotFound)
			return
		}

		utils.WriteJSON(w, 200, response.Teams[0])
	}
}
