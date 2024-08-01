package impl

import (
	"final-project-enigma/config"
	"final-project-enigma/entity"
	"fmt"
)

type WorkRepository struct{}

func NewWorkRepository() *WorkRepository {
	return &WorkRepository{}
}

func (WorkRepository) CreateWork(work entity.Work) (entity.Work, error) {
	if result := config.DB.Create(&work); result.Error != nil {
		return entity.Work{}, result.Error
	}
	return work, nil
}

func (WorkRepository) UpdateWork(work entity.Work) (entity.Work, error) {
	if result := config.DB.Save(&work); result.Error != nil {
		return entity.Work{}, result.Error
	}
	return work, nil
}

func (WorkRepository) DeleteWork(id string) error {
	if result := config.DB.Delete(&entity.Work{}, "id = ?", id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (WorkRepository) GetById(id string) (entity.Work, error) {
	var work entity.Work
	config.DB.Where("id = ?", id).First(&work)
	if work.ID == "" {
		return entity.Work{}, fmt.Errorf("data not found")
	}
	return work, nil
}
