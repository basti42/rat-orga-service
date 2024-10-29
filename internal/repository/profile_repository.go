package repository

import (
	"github.com/basti42/orga-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (pr *ProfileRepository) AddNewProfile(profile *models.Profile) error {
	if tx := pr.db.Create(profile); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (pr *ProfileRepository) GetProfileForUser(userUUID uuid.UUID) (*models.Profile, error) {
	var profile models.Profile
	if tx := pr.db.Where("user_uuid = ?", userUUID).First(&profile); tx.Error != nil {
		return nil, tx.Error
	}
	return &profile, nil
}
