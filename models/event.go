package models

type Event struct {
	EventID       string   `json:"eventId"`
	Code          string   `json:"code"`
	DivisionCode  string   `json:"divisionCode"`
	Name          string   `json:"name"`
	Remote        bool     `json:"remote"`
	Hybrid        bool     `json:"hybrid"`
	FieldCount    int      `json:"fieldCount"`
	Address       string   `json:"address"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	Country       string   `json:"country"`
	Website       string   `json:"website"`
	LiveStreamURL string   `json:"liveStreamUrl"`
	Webcasts      []string `json:"webcasts"`
	DateStart     string   `json:"dateStart"`
	DateEnd       string   `json:"dateEnd"`
}

