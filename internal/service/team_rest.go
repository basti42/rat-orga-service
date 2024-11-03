package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/basti42/orga-service/internal/models"
	"github.com/basti42/orga-service/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamService struct {
	repo *repository.TeamRepository
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{repo: repository.NewTeamRepository(db)}
}

func (ts *TeamService) HandleCreateDefaultTeam(c *gin.Context) (*models.Team, error) {
	userString, ok := c.Keys["user-uuid"].(string)
	if !ok {
		return nil, errors.New("no user found in context from token verification")
	}
	userUUID, err := uuid.Parse(userString)
	if err != nil {
		return nil, errors.New("malformated user uuid from token")
	}
	profile := c.Keys["profile"].(models.Profile)
	newUUID, _ := uuid.NewRandom()
	defaultTeam := models.Team{
		UUID:         newUUID,
		OwnerUUID:    userUUID,
		Abbreviation: strings.ToUpper(fmt.Sprintf("%v", profile.Name[0])) + "AT",
		Name:         fmt.Sprintf("%v's Team", profile.Name),
		Members:      []*models.Profile{&profile},
	}
	if err := ts.repo.AddNewTeam(&defaultTeam); err != nil {
		return nil, err
	}
	return &defaultTeam, nil
}

func (ts *TeamService) HandleAddNewTeam(c *gin.Context) (*models.Team, error) {
	userString, ok := c.Keys["user-uuid"].(string)
	if !ok {
		return nil, errors.New("no user found in context from token verification")
	}
	userUUID, err := uuid.Parse(userString)
	if err != nil {
		return nil, errors.New("malformated user uuid from token")
	}
	var newTeamRequest models.NewTeamRequest
	if err := c.BindJSON(&newTeamRequest); err != nil {
		return nil, errors.New("error binding new team request")
	}

	teamUUID, _ := uuid.NewRandom()
	newTeam := models.Team{
		UUID:         teamUUID,
		OwnerUUID:    userUUID,
		Name:         newTeamRequest.Name,
		Abbreviation: newTeamRequest.Abbreviation,
	}

	if err = ts.repo.AddNewTeam(&newTeam); err != nil {
		return nil, err
	}
	return &newTeam, nil
}

func (ts *TeamService) HandleGetUserTeams(c *gin.Context) ([]models.Team, error) {
	userString, ok := c.Keys["user-uuid"].(string)
	if !ok {
		return nil, errors.New("no user found in context from token verification")
	}
	userUUID, err := uuid.Parse(userString)
	if err != nil {
		return nil, errors.New("malformated user uuid from token")
	}
	return ts.repo.GetTeamsForUser(userUUID)
}

func (ts *TeamService) HandleGetPublicProfiles(c *gin.Context) ([]models.PublicProfile, error) {
	userString, ok := c.Keys["user-uuid"].(string)
	if !ok {
		return nil, errors.New("no user found in context from token verification")
	}
	userUUID, err := uuid.Parse(userString)
	if err != nil {
		return nil, errors.New("malformated user uuid from token")
	}

	teamUUIDString := c.Param("team-uuid")
	teamUUID, err := uuid.Parse(teamUUIDString)
	if err != nil {
		return nil, errors.New("malformatted team uuid in path")
	}

	profiles, err := ts.repo.GetTeamMemberProfiles(teamUUID, userUUID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error fetching mambers for team: %v", err))
	}

	// convert to PublicProfile
	publicProfiles := make([]models.PublicProfile, len(profiles))
	for i, profile := range profiles {
		publicProfiles[i] = models.PublicProfile{
			UserUUID:     profile.UUID,
			Name:         profile.Name,
			Abbreviation: profile.Abbreviation,
			AvatarURL:    profile.AvatarURL,
		}
	}

	return publicProfiles, nil
}
