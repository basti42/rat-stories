package main

import (
	"fmt"
	"log"

	application "github.com/basti42/stories-service/internal"
	"github.com/basti42/stories-service/internal/middlewares"
	"github.com/basti42/stories-service/internal/repository"
	"github.com/basti42/stories-service/internal/system"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db := repository.GetDatabaseConnection()
	app := application.NewApplication(db)

	router.GET("/health", app.Health)

	// TODO middleware to handle access, once auth is in place

	storiesGroup := router.Group("/rat/stories")
	storiesGroup.Use(middlewares.UserValidationMiddleware())

	storiesGroup.GET("", app.ListStories)
	storiesGroup.POST("", app.AddNewStory)
	storiesGroup.GET("/:story-uuid", app.GetStory)
	storiesGroup.POST("/:story-uuid/status", app.UpdateStatusHistory)
	storiesGroup.POST("/:story-uuid/acceptance-criteria", app.AddAcceptanceCriterium)

	port := system.PORT
	serviceName := system.SERVICE_NAME

	log.Printf("starting [%v] on port=%v", serviceName, port)
	if err := router.Run(fmt.Sprintf(":%v", port)); err != nil {
		log.Panicf("error starting [%v] on port=%v: %v", serviceName, port, err)
	}

}
