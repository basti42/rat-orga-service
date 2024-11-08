package models

import (
	"github.com/google/uuid"
)

type NewTeamRequest struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type TeamAPI struct {
	UUID         uuid.UUID       `gorm:"primaryKey" json:"uuid"`
	OwnerUUID    uuid.UUID       `json:"owner_uuid"`
	Abbreviation string          `json:"abbreviation"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Members      []TeamMemberAPI `json:"members"`
}

type TeamMemberAPI struct {
	TeamUUID uuid.UUID `json:"team_uuid"`
	UserUUID uuid.UUID `json:"user_uuid"`
	Role     string    `json:"role"`
}
