package impl

import (
	"errors"
	"final-project-enigma/config"
	"final-project-enigma/entity"
	"final-project-enigma/helper"
	"strconv"
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
		return entity.Work{}, errors.New("data not found")
	}
	return work, nil
}

func (WorkRepository) GetAllWork(page, size string) ([]entity.Work, string, error) {
	var works []entity.Work
	config.DB.Scopes(helper.Paginate(page, size)).Find(&works)
	if len(works) == 0 {
		return nil, "0", errors.New("data not found")
	}
	var total int64
	config.DB.Model(&entity.Work{}).Count(&total)
	return works, strconv.FormatInt(total, 10), nil
}
