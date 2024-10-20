package repository

import (
	"fmt"
	"log"

	"github.com/basti42/stories-service/internal/models"
	"github.com/basti42/stories-service/internal/system"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		system.DB_SERVER, system.DATABASE_USER, system.DATABASE_PASSWORD, system.DATABASE, system.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("error creating stories table: %v", err)
	}

	if err = db.AutoMigrate(&models.Story{}, &models.Estimation{}, &models.AcceptanceCriterium{}, &models.StoryHistory{}); err != nil {
		log.Panicf("error migrating stories DB: %v", err)
	}

	return db
}
