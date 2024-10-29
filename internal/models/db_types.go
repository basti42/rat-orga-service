package models

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	UUID         uuid.UUID  `gorm:"primaryKey" json:"uuid"`
	OwnerUUID    uuid.UUID  `json:"owner_uuid"`
	Abbreviation string     `json:"abbreviation"`
	Name         string     `json:"name"`
	Members      []*Profile `gorm:"many2many:teams_profiles" json:"members"`
}

func (Team) TableName() string { return "teams" }

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
	Teams        []*Team   `gorm:"many2many:teams_profiles" json:"teams"`
}

func (Profile) TableName() string { return "profiles" }
