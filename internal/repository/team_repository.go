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

func (tr *TeamRepository) AddNewTeamMember(member *models.TeamMember) error {
	if tx := tr.db.Create(member); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (tr *TeamRepository) GetTeamsForUser(userUUID uuid.UUID) ([]models.TeamAPI, error) {

	// 1. find all teams for that user, should be only one in the beginning
	var teams []models.Team
	tx := tr.db.Where("owner_uuid = ?", userUUID).Find(&teams)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// 2. iterate all found teams to retrieve their members and prepare the output api type
	teamsAPI := make([]models.TeamAPI, len(teams))
	for i, team := range teams {

		var members []models.TeamMember
		if tx := tr.db.Where("team_uuid = ?", team.UUID).Find(&members); tx.Error != nil {
			return nil, tx.Error
		}

		membersAPI := make([]models.TeamMemberAPI, len(members))
		for m, member := range members {
			membersAPI[m] = models.TeamMemberAPI{
				TeamUUID: member.TeamUUID,
				UserUUID: member.UserUUID,
				Role:     member.Role,
			}
		}

		teamsAPI[i] = models.TeamAPI{
			UUID:         team.UUID,
			OwnerUUID:    team.OwnerUUID,
			Name:         team.Name,
			Abbreviation: team.Abbreviation,
			Description:  team.Description,
			Members:      membersAPI,
		}
	}
	return teamsAPI, nil
}

func (tr *TeamRepository) GetTeamMembers(teamUUID uuid.UUID) ([]*models.TeamMember, error) {
	var members []*models.TeamMember
	if tx := tr.db.
		Where("team_uuid = ?", teamUUID).
		First(&members); tx.Error != nil {
		return nil, tx.Error
	}
	return members, nil
}
