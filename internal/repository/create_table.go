package repository

import (
	"io/fs"
	"log"
	"os"

	"github.com/basti42/stories-service/internal/models"
	"github.com/basti42/stories-service/internal/system"
	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	// TODO find better permission bits
	if err := os.MkdirAll(system.DB_PATH, fs.ModePerm); err != nil {
		log.Panicf("error creating stories db directory: %v", err)
	}

	db, err := gorm.Open(sqlite.Open("./data/stories.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("error creating stories table: %v", err)
	}

	if err = db.AutoMigrate(&models.Story{}, &models.Estimation{}, &models.AcceptanceCriterium{}, &models.StoryHistory{}); err != nil {
		log.Panicf("error migrating stories DB: %v", err)
	}

	return db
}
