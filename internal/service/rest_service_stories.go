package service

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/basti42/stories-service/internal/models"
	"github.com/basti42/stories-service/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestService struct {
	repo *repository.StoriesRepository
}

func NewRestService(db *gorm.DB) *RestService {
	return &RestService{repo: repository.NewStoriesRepository(db)}
}

func (svc *RestService) HandleListStories(c *gin.Context) ([]models.Story, error) {
	userUUID, err := uuid.Parse(c.Keys["user-uuid"].(string))
	if err != nil {
		return nil, errors.New("no valid user from request")
	}
	limitQuery := c.Query("limit")
	var limit int = 100
	if limitQuery != "" {
		limitValue, err := strconv.ParseInt(limitQuery, 10, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error extracting query paramter 'limit'=%v", limitQuery))
		}
		limit = int(limitValue)
	}
	if limit < 0 {
		return nil, errors.New("'limit' < 0, does not work")
	}

	offsetQuery := c.Query("offset")
	var offset int = 0
	if offsetQuery != "" {
		offsetValue, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error extracting query parameter 'offset'=%v", offsetQuery))
		}
		offset = int(offsetValue)
	}
	if offset < 0 {
		return nil, errors.New("'offset' < 0, does not work")
	}

	slog.Info(fmt.Sprintf("limit = %v, offset = %v", limit, offset))
	slog.Info(fmt.Sprintf("user uuid = %v", userUUID))

	return svc.repo.ListStories(userUUID, offset, limit)

}

func (svc *RestService) HandleAddNewStory(c *gin.Context) (models.Story, error) {
	var newStory models.NewStory
	if err := c.BindJSON(&newStory); err != nil {
		slog.Error("error binding JSON POST request: ", err)
		return models.Story{}, err
	}

	userUUID, err := uuid.Parse(c.Keys["user-uuid"].(string))
	if err != nil {
		return models.Story{}, errors.New("no valid user from request")
	}

	storyUUID, _ := uuid.NewRandom()
	historyUUID, _ := uuid.NewRandom()
	now := time.Now().UTC()

	story := &models.Story{
		UUID:               storyUUID,
		Creator:            userUUID,
		Assignee:           newStory.Assignee,
		Team:               newStory.Team,
		Project:            newStory.Project,
		Feature:            newStory.Feature,
		CreatedAt:          now,
		UpdatedAt:          now,
		ClosedAt:           nil,
		Status:             newStory.Status,
		Title:              newStory.Title,
		Description:        newStory.Description,
		AcceptanceCriteria: newStory.AcceptanceCriteria,
		History: []models.StoryHistory{
			{
				UUID:       historyUUID,
				StoryUUID:  storyUUID,
				FromStatus: "",
				ToStatus:   "icebox",
				Date:       now,
				By:         userUUID,
			},
		},
	}

	story, err = svc.repo.AddNewStory(story)
	if err != nil {
		slog.Error("error adding story to db: ", err)
		return models.Story{}, err
	}

	return *story, nil
}

func (svc *RestService) HandleGetStory(c *gin.Context) (*models.Story, error) {
	storyUUID, err := uuid.Parse(c.Param("story-uuid"))
	if err != nil {
		return nil, errors.New("missing parameter 'story-uuid' in request")
	}
	userUUID, err := uuid.Parse(c.Keys["user-uuid"].(string))
	if err != nil {
		return nil, errors.New("no valid user from request")
	}
	return svc.repo.GetStory(storyUUID, userUUID)
}
