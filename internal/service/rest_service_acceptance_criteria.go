package service

import (
	"errors"

	"github.com/basti42/stories-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (svc *RestService) HandleAddAcceptanceCriterium(c *gin.Context) (*models.Story, error) {
	storyUUID, err := uuid.Parse(c.Param("story-uuid"))
	if err != nil {
		return nil, errors.New("missing parameter 'story-uuid' in request")
	}

	userUUID, err := uuid.Parse(c.Keys["user-uuid"].(string))
	if err != nil {
		return nil, errors.New("no valid user from request")
	}

	var newAC models.NewAcceptanceCriterium
	if err := c.BindJSON(&newAC); err != nil {
		return nil, errors.New("error unmarshalling new acceptance criterium request body")
	}

	if newAC.Description == "" {
		return nil, errors.New("empty acceptance criterium")
	}

	newUUID, _ := uuid.NewRandom()
	ac := models.AcceptanceCriterium{
		UUID:        newUUID,
		StoryUUID:   storyUUID,
		Description: newAC.Description,
		AcceptedAt:  nil,
		AcceptedBy:  nil,
	}

	if err := svc.repo.AddAcceptanceCriterium(ac); err != nil {
		return nil, err
	}

	return svc.repo.GetStory(storyUUID, userUUID)
}
