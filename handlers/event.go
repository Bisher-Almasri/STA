package handlers

import (
	"STA/models"
	"STA/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// @Summary Get event
// @Description Returns a event by code
// @Tags Events
// @Produce json
// @Param number path string true "Event Code"
// @Success 200 {object} models.EventResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /events/{number} [get]
func GetEvent(client http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const prefix = "/api/v1/event/"
		if !strings.HasPrefix(r.URL.Path, prefix) {
			http.NotFound(w, r)
			return
		}

		eventStr := strings.TrimPrefix(r.URL.Path, prefix)

		url := fmt.Sprintf("events?eventCode=%s", eventStr)

		bodyBytes, err := utils.ApiRequest(url, client)
		if err != nil {
			log.Printf("ApiRequest error: %v", err)
			utils.WriteError(w, "upstream API error", http.StatusBadGateway)
			return
		}

		var response models.EventResponse
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			log.Printf("decode error: %v", err)
			utils.WriteError(w, "failed to decode team response", http.StatusBadGateway)
			return
		}

		if len(response.Events) == 0 {
			utils.WriteError(w, "event not found", http.StatusNotFound)
			return
		}

		utils.WriteJSON(w, 200, response.Events[0])
	}
}
