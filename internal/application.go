package internal

import (
	"fmt"
	"net/http"

	"github.com/basti42/orga-service/internal/models"
	"github.com/basti42/orga-service/internal/service"
	"github.com/gin-gonic/gin"
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
	team, err := service.NewTeamService(app.db).HandleCreateDefaultTeam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("error creating team for user: %v", err.Error())})
	}
	profile.Teams = []*models.Team{team}
	c.JSON(http.StatusOK, profile)
	return
}

func (app *Application) GetPublicProfiles(c *gin.Context) {
	publicProfiles, err := service.NewTeamService(app.db).HandleGetPublicProfiles(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, publicProfiles)
}

/*
SPRINT
*/
func (app *Application) AddNewSprint(c *gin.Context) {
	// TODO
	return
}
