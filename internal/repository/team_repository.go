package repository

import (
	"github.com/basti42/orga-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (tr *TeamRepository) AddNewTeam(team *models.Team) error {
	if tx := tr.db.Create(team); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (tr *TeamRepository) GetTeamsForUser(userUUID uuid.UUID) ([]models.Team, error) {
	var teams []models.Team
	tx := tr.db.Preload("Members").Where("owner_uuid = ?", userUUID).Find(&teams)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return teams, nil
}

func (tr *TeamRepository) GetTeamMemberProfiles(teamUUID, userUUID uuid.UUID) ([]*models.Profile, error) {
	var team models.Team
	if tx := tr.db.
		Preload("Members").
		Where("uuid = ? AND owner_uuid = ?", teamUUID, userUUID).
		First(&team); tx.Error != nil {
		return nil, tx.Error
	}
	return team.Members, nil
}
