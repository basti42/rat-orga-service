package models

import "github.com/google/uuid"

type NewProfileRequest struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	AvatarURL    string `json:"avatar_url"`
}

type PublicProfile struct {
	UserUUID     uuid.UUID `json:"user_uuid"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	AvatarURL    string    `json:"avatar_url"`
}

type ProfileUpdateRequest struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Bio          string `json:"bio"`
	Quote        string `json:"quote"`
}
