package models

type Team struct {
	TeamNumber        int     `json:"teamNumber"`
	DisplayTeamNumber *string `json:"displayTeamNumber"`
	TeamID            int     `json:"teamId"`
	TeamProfileID     int     `json:"teamProfileId"`
	NameFull          *string `json:"nameFull"`
	NameShort         *string `json:"nameShort"`
	SchoolName        *string `json:"schoolName"`
	City              *string `json:"city"`
	StateProv         *string `json:"stateProv"`
	Country           *string `json:"country"`
	Website           *string `json:"website"`
	RookieYear        *int    `json:"rookieYear"`
	RobotName         *string `json:"robotName"`
	DistrictCode      *string `json:"districtCode"`
	HomeCMP           *string `json:"homeCMP"`
	HomeRegion        *string `json:"homeRegion"`
	DisplayLocation   *string `json:"displayLocation"`
}

type TeamResponse struct {
	Teams []Team `json:"teams"`
}
