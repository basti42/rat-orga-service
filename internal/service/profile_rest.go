package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/basti42/orga-service/internal/models"
	"github.com/basti42/orga-service/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ProfileService struct {
	profileRepo *repository.ProfileRepository
}

func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{profileRepo: repository.NewProfileRepository(db)}
}

func (ps *ProfileService) HandleAddNewProfile(c *gin.Context) (models.Profile, error) {
	userString, ok := c.Keys["user-uuid"]
	if !ok {
		return models.Profile{}, errors.New("no user found in context from token validation")
	}
	userUUID, err := uuid.Parse(fmt.Sprintf("%v", userString))
	if err != nil {
		return models.Profile{}, errors.New(fmt.Sprintf("malformatted user uuid found in token: %v", userString))
	}

	var newProfileRequest models.NewProfileRequest
	if err := c.BindJSON(&newProfileRequest); err != nil {
		return models.Profile{}, errors.New(fmt.Sprintf("error binding request body: %v", err))
	}

	newUUID, _ := uuid.NewRandom()
	now := time.Now()
	profile := models.Profile{
		UUID:         newUUID,
		UserUUID:     userUUID,
		CreatedAt:    now,
		UpdatedAt:    now,
		Name:         newProfileRequest.Name,
		Abbreviation: newProfileRequest.Abbreviation,
		AvatarURL:    newProfileRequest.AvatarURL,
		Role:         "",
		Bio:          "",
		Quote:        "",
	}
	if err := ps.profileRepo.AddNewProfile(&profile); err != nil {
		return models.Profile{}, errors.New(fmt.Sprintf("error creating profile in db: %v", err))
	}

	return profile, nil
}
