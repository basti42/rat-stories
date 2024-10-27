package repository

import (
	"fmt"
	"log/slog"

	"github.com/basti42/stories-service/internal/models"
	"github.com/google/uuid"
)

func (repo *StoriesRepository) ListStories(userUUID uuid.UUID, offset, limit int) ([]models.Story, error) {
	var stories []models.Story
	tx := repo.db.
		Preload("AcceptanceCriteria").
		Preload("History").
		Preload("Comments").
		Where("creator = ?", userUUID).
		Order("created_at desc").
		Offset(offset).
		Limit(limit).
		Find(&stories)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return stories, nil
}

func (repo *StoriesRepository) AddNewStory(story *models.Story) (*models.Story, error) {
	createStoryResult := repo.db.Create(story)
	if createStoryResult.Error != nil {
		return nil, createStoryResult.Error
	}

	newStory, err := repo.GetStory(story.UUID, story.Creator)
	if err != nil {
		return nil, err
	}
	return newStory, nil
}

func (repo *StoriesRepository) GetStory(storyUUID, userUUID uuid.UUID) (*models.Story, error) {
	var story models.Story
	if tx := repo.db.
		Preload("AcceptanceCriteria").
		Preload("History").
		Preload("Comments").
		Where("uuid = ? AND creator = ?", storyUUID, userUUID).
		First(&story); tx.Error != nil {
		slog.Warn(fmt.Sprintf("error retrieving story=%v from database: %v", storyUUID, tx.Error))
		return nil, tx.Error
	}
	return &story, nil
}

func (repo *StoriesRepository) UpdateStoryStatus(storyUUID, userUUID uuid.UUID, status string) error {
	story, err := repo.GetStory(storyUUID, userUUID)
	if err != nil {
		return err
	}
	story.Status = status
	if tx := repo.db.Save(story); tx.Error != nil {
		return tx.Error
	}
	return nil
}
