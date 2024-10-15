package application

import (
	"net/http"

	"github.com/basti42/stories-service/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	db *gorm.DB
}

func NewApplication(db *gorm.DB) *Application {
	return &Application{db: db}
}

func (a *Application) Health(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (a *Application) ListStories(c *gin.Context) {
	stories, err := service.NewRestService(a.db).HandleListStories(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, stories)
}

func (a *Application) AddNewStory(c *gin.Context) {
	newStory, err := service.NewRestService(a.db).HandleAddNewStory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, newStory)
}

func (a *Application) GetStory(c *gin.Context) {
	story, err := service.NewRestService(a.db).HandleGetStory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, story)
}

func (a *Application) UpdateStatusHistory(c *gin.Context) {
	story, err := service.NewRestService(a.db).HandleUpdateStoryStatusHistory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, story)
}

func (a *Application) AddAcceptanceCriterium(c *gin.Context) {
	story, err := service.NewRestService(a.db).HandleAddAcceptanceCriterium(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, story)
}
