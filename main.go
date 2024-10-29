package main

import (
	"fmt"
	"log"

	"github.com/basti42/orga-service/internal"
	"github.com/basti42/orga-service/internal/middlewares"
	"github.com/basti42/orga-service/internal/repository"
	"github.com/basti42/orga-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	serviceName := utils.SERVICE_NAME
	log.SetPrefix(fmt.Sprintf("[%v] ", serviceName))

	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
		gin.Recovery(),
	)

	db := repository.GetDatabaseConnection()
	app := internal.NewApplication(db)

	router.GET("/health", app.Health)

	teamGroup := router.Group("/rat/teams")
	teamGroup.Use(middlewares.UserValidationMiddleware())
	teamGroup.POST("", app.AddNewTeam)
	teamGroup.GET("/:team-uuid/public-profiles", app.GetPublicProfiles)

	profileGroup := router.Group("/rat/profiles")
	profileGroup.Use(middlewares.UserValidationMiddleware())
	profileGroup.POST("", app.AddNewProfile)

	sprintGroup := router.Group("/rat/sprints")
	sprintGroup.Use(middlewares.UserValidationMiddleware())
	sprintGroup.POST("", app.AddNewSprint)

	if err := router.Run(fmt.Sprintf(":%v", utils.PORT)); err != nil {
		log.Panicf("error starting [%v] on port=%v", utils.SERVICE_NAME, utils.PORT)
	}
}
