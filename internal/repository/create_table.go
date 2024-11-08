package repository

import (
	"fmt"
	"log"

	"github.com/basti42/orga-service/internal/models"
	"github.com/basti42/orga-service/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.DB_HOST, utils.DB_USER, utils.DB_PASSWORD, utils.DB_NAME, utils.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("error creating stories table: %v", err)
	}

	if err = db.AutoMigrate(&models.Profile{}, &models.Team{}, &models.TeamMember{}); err != nil {
		log.Panicf("error migrating orga DB: %v", err)
	}

	return db

}
