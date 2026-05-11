package models

type Team struct {
	TeamNumber int    `json:"teamNumber"`
	Name       string `json:"name"`
	SchoolName string `json:"schoolName"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	Website    string `json:"website"`
	RookieYear int    `json:"rookieYear"`
	RobotName  string `json:"robotName"`
	HomeRegion string `json:"homeRegion"`
}

