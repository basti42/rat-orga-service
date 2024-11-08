package internal

import (
	"fmt"
	"net/http"

	"github.com/basti42/orga-service/internal/models"
	"github.com/basti42/orga-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	db *gorm.DB
}

func NewApplication(db *gorm.DB) *Application {
	return &Application{db: db}
}

func (app *Application) Health(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
	return
}

/*
TEAM
*/
func (app *Application) AddNewTeam(c *gin.Context) {
	team, err := service.NewTeamService(app.db).HandleAddNewTeam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, team)
	return
}

func (app *Application) GetUserTeams(c *gin.Context) {
	teams, err := service.NewTeamService(app.db).HandleGetUserTeams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

/*
PROFILE
*/
func (app *Application) AddNewProfile(c *gin.Context) {
	profile, err := service.NewProfileService(app.db).HandleAddNewProfile(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("error creating profile for user: %v", err.Error())})
		return
	}
	// once a profile was created successfully create a team also
	c.Set("profile", profile)
	_, err = service.NewTeamService(app.db).HandleCreateDefaultTeam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("error creating team for user: %v", err.Error())})
	}
	c.JSON(http.StatusOK, profile)
	return
}

func (app *Application) GetUserProfile(c *gin.Context) {
	profile, err := service.NewProfileService(app.db).HandleGetUserProfile(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mesage": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (app *Application) UpdateUserProfile(c *gin.Context) {
	profile, err := service.NewProfileService(app.db).HandleUpdateProfile(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (app *Application) GetPublicProfiles(c *gin.Context) {
	// 1. get teammates from teams repository
	teamMembers, err := service.NewTeamService(app.db).HandleGetTeamMembers(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 2. then get profiles for each member of the team
	userUUIDs := make([]uuid.UUID, len(teamMembers))
	for i, teamMember := range teamMembers {
		userUUIDs[i] = teamMember.UserUUID
	}
	c.Set("team_member_uuids", userUUIDs)

	publicProfiles, err := service.NewProfileService(app.db).HandleGetPublicProfiles(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 3. role is team specific set it to the profile information
	// map [userUUID] -> teamMember
	members := make(map[uuid.UUID]models.TeamMember, len(teamMembers))
	for _, tm := range teamMembers {
		members[tm.UserUUID] = *tm
	}

	for i, profile := range publicProfiles {
		publicProfiles[i].Role = members[profile.UserUUID].Role
	}

	// 4. if all ok return the found public user profiles
	c.JSON(http.StatusOK, publicProfiles)
}

/*
SPRINT
*/
func (app *Application) AddNewSprint(c *gin.Context) {
	// TODO
	return
}
