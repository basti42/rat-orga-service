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

func (pr *ProfileRepository) UpdateProfile(profileUUID, userUUID uuid.UUID, updates models.ProfileUpdateRequest) (*models.Profile, error) {
	profile, err := pr.GetProfileForUser(userUUID)
	if err != nil {
		return nil, err
	}
	profile.Name = updates.Name
	profile.Abbreviation = updates.Abbreviation
	profile.Bio = updates.Bio
	profile.Quote = updates.Quote
	if tx := pr.db.Save(profile); tx.Error != nil {
		return nil, tx.Error
	}
	return profile, nil
}

func (pr *ProfileRepository) GetProfiles(userUUIDs []uuid.UUID) ([]models.Profile, error) {
	var profiles []models.Profile
	if tx := pr.db.
		Model(&models.Profile{}).
		Where("user_uuid IN ?", userUUIDs).
		Find(&profiles); tx.Error != nil {
		return nil, tx.Error
	}
	return profiles, nil
}
