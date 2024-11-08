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

func (ts *TeamService) HandleGetUserTeams(c *gin.Context) ([]models.TeamAPI, error) {
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

func (ts *TeamService) HandleGetTeamMembers(c *gin.Context) ([]*models.TeamMember, error) {
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

	teamMembers, err := ts.repo.GetTeamMembers(teamUUID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error fetching team members for team: %v", err))
	}

	// check if user is in the team members for this team, else do not return
	requestingUserIsTeamMember := false
	for _, teamMember := range teamMembers {
		if teamMember.UserUUID == userUUID {
			requestingUserIsTeamMember = true
			break
		}
	}

	if !requestingUserIsTeamMember {
		msg := fmt.Sprintf("user=%v is not a member of team=%v", userUUID, teamUUID)
		return nil, errors.New(msg)
	}

	return teamMembers, nil
}
