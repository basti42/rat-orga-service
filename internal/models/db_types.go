package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	UUID         uuid.UUID `gorm:"primaryKey" json:"uuid"`
	OwnerUUID    uuid.UUID `json:"owner_uuid"`
	Abbreviation string    `json:"abbreviation"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
}

func (Team) TableName() string { return "teams" }

type TeamMember struct {
	gorm.Model
	TeamUUID uuid.UUID `json:"team_uuid"`
	UserUUID uuid.UUID `json:"user_uuid"`
	Role     string    `json:"role"`
}

func (TeamMember) TableName() string { return "team_members" }

type Profile struct {
	UUID         uuid.UUID `gorm:"primaryKey" json:"uuid"`
	UserUUID     uuid.UUID `gorm:"unique" json:"user_uuid"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	AvatarURL    string    `json:"avatar_url"`
	Role         string    `json:"role"`
	Bio          string    `json:"bio"`
	Quote        string    `json:"quote"`
}

func (Profile) TableName() string { return "profiles" }
