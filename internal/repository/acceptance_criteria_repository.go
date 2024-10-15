package repository

import "github.com/basti42/stories-service/internal/models"

func (repo *StoriesRepository) AddAcceptanceCriterium(ac models.AcceptanceCriterium) error {
	if tx := repo.db.Create(ac); tx.Error != nil {
		return tx.Error
	}
	return nil
}
