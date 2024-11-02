package models

type NewTeamRequest struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}
