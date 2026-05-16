package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const BASE_URL = "https://ftc-api.firstinspires.org/v2.0/2025"

func ApiRequest(path string, client http.Client) ([]byte, error) {
	targetURL := fmt.Sprintf("%s/%s", BASE_URL, path)

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "STA")

	req.SetBasicAuth(
		os.Getenv("FTC_EVENTS_USERNAME"),
		os.Getenv("FTC_EVENTS_AUTH_TOKEN"),
	)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("FTC API returned %d: %s", res.StatusCode, string(body))
	}

	return body, nil
}
