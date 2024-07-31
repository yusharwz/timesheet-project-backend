package repository

import (
	"final-project-enigma/entity"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type TimeSheetRepository interface {
	CreateTimeSheet(ts *entity.TimeSheet) (*entity.TimeSheet, error)
	FindByIdTimeSheet(id string) (*entity.TimeSheet, error)
	FindAllTimeSheet() (*[]entity.TimeSheet, error)
	UpdateTimeSheet(ts *entity.TimeSheet) (*entity.TimeSheet, error)
	DeleteTimeSheet(id string) error
}

type timeSheetRepository struct {
	db *gorm.DB
}

func (t *timeSheetRepository) CreateTimeSheet(ts *entity.TimeSheet) (*entity.TimeSheet, error) {
	if err := t.db.Create(ts).Error; err != nil {
		log.Error().Err(err).Msg("Failed to create timesheet")
		return nil, err
	}
	return ts, nil
}
func (t *timeSheetRepository) FindByIdTimeSheet(id string) (*entity.TimeSheet, error) {
	var ts entity.TimeSheet
	if err := t.db.First(&ts, "id = ? AND deleted_at IS NULL", id).Error; err != nil {
		log.Error().Err(err).Msgf("TimeSheet with id %s not found", id)
		return nil, err
	}

	return &ts, nil
}
func (t *timeSheetRepository) FindAllTimeSheet() (*[]entity.TimeSheet, error) {
	var ts []entity.TimeSheet
	if err := t.db.Where("deleted_at IS NULL").Find(&ts).Error; err != nil {
		log.Error().Err(err).Msg("data is null")
		return nil, err
	}

	return &ts, nil

}
func (t *timeSheetRepository) UpdateTimeSheet(ts *entity.TimeSheet) (*entity.TimeSheet, error) {
	if err := t.db.Save(ts).Error; err != nil {
		log.Error().Err(err).Msg("failed to update timesheet")
		return nil, err
	}
	return ts, nil
}
func (t *timeSheetRepository) DeleteTimeSheet(id string) error {
	var ts entity.TimeSheet
	if err := t.db.First(&ts, "id = ? AND deleted_at IS NULL", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil // Tidak ada error jika record tidak ditemukan
		}
		return err
	}

	// Soft delete: set deleted_at
	if err := t.db.Model(&ts).Update("deleted_at", gorm.DeletedAt{Time: time.Now(), Valid: true}).Error; err != nil {
		return err
	}
	return nil
}


func NewTimeSheetRepository(db *gorm.DB) TimeSheetRepository {
	return &timeSheetRepository{db: db}
}