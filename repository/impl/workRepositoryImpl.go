package impl

import (
	"errors"
	"final-project-enigma/config"
	"final-project-enigma/entity"
	"final-project-enigma/helper"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type WorkRepository struct{}

func NewWorkRepository() *WorkRepository {
	return &WorkRepository{}
}

func (WorkRepository) CreateWork(work entity.Work) (entity.Work, error) {
	if result := config.DB.Create(&work); result.Error != nil {
		log.Error()
		return entity.Work{}, result.Error
	}
	return work, nil
}

func (WorkRepository) UpdateWork(work entity.Work) (entity.Work, error) {
	if result := config.DB.Save(&work); result.Error != nil {
		log.Error()
		return entity.Work{}, result.Error
	}
	return work, nil
}

func (WorkRepository) DeleteWork(id string) error {
	if result := config.DB.Delete(&entity.Work{}, "id = ?", id); result.Error != nil {
		log.Error()
		return result.Error
	}
	return nil
}

func (WorkRepository) GetById(id string) (entity.Work, error) {
	var work entity.Work
	config.DB.Where("id = ?", id).First(&work)
	if work.ID == "" {
		log.Error()
		return entity.Work{}, errors.New("data not found")
	}
	return work, nil
}

func (WorkRepository) GetAllWork(spec []func(db *gorm.DB) *gorm.DB) ([]entity.Work, string, error) {
	var works []entity.Work
	db := config.DB.Scopes(spec...).Find(&works)
	totalRows := helper.GetTotalRows(db)
	return works, totalRows, db.Error
}
