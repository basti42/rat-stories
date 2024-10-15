package service

import (
	"errors"
	"time"

	"github.com/basti42/stories-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (svc *RestService) HandleUpdateStoryStatusHistory(c *gin.Context) (*models.Story, error) {
	storyUUID, err := uuid.Parse(c.Param("story-uuid"))
	if err != nil {
		return nil, errors.New("missing parameter 'story-uuid' in request")
	}

	userUUID, err := uuid.Parse(c.Keys["user-uuid"].(string))
	if err != nil {
		return nil, errors.New("no valid user from request")
	}

	var newStatusHistory models.NewStoryHistory
	if err = c.BindJSON(&newStatusHistory); err != nil {
		return nil, errors.New("error unmarshalling request body")
	}

	if newStatusHistory.FromStatus == "" && newStatusHistory.ToStatus == "" {
		return nil, errors.New("bad request, no status transition provided")
	}

	newUUID, _ := uuid.NewRandom()

	history := models.StoryHistory{
		UUID:       newUUID,
		StoryUUID:  storyUUID,
		FromStatus: newStatusHistory.FromStatus,
		ToStatus:   newStatusHistory.ToStatus,
		Date:       time.Now().UTC(),
		By:         uuid.MustParse(c.Keys["user-uuid"].(string)),
	}

	if err = svc.repo.AddStatusHistory(history); err != nil {
		return nil, err
	}
	if err = svc.repo.UpdateStoryStatus(storyUUID, userUUID, history.ToStatus); err != nil {
		return nil, err
	}
	return svc.repo.GetStory(storyUUID, userUUID)

}
