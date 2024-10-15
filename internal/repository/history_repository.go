package repository

import "github.com/basti42/stories-service/internal/models"

func (repo *StoriesRepository) AddStatusHistory(history models.StoryHistory) error {
	if tx := repo.db.Create(&history); tx.Error != nil {
		return tx.Error
	}
	return nil
}
