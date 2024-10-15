package repository

import "gorm.io/gorm"

type StoriesRepository struct {
	db *gorm.DB
}

func NewStoriesRepository(db *gorm.DB) *StoriesRepository {
	return &StoriesRepository{db: db}
}
