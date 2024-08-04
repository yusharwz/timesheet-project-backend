package impl

import (
	"final-project-enigma/config"
	"final-project-enigma/entity"
	"gorm.io/gorm"
	"time"
)

type TimeSheetRepository struct{}

func NewTimeSheetRepository() *TimeSheetRepository {
	return &TimeSheetRepository{}
}

func (TimeSheetRepository) CreateTimeSheet(ts entity.TimeSheet) (entity.TimeSheet, error) {
	err := config.DB.Create(&ts).Error
	if err != nil {
		return entity.TimeSheet{}, err
	}
	return ts, nil
}

func (TimeSheetRepository) UpdateTimeSheet(ts *entity.TimeSheet) error {
	return config.DB.Save(ts).Error
}

func (TimeSheetRepository) DeleteTimeSheet(id string) error {
	return config.DB.Model(&entity.TimeSheet{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
}

func (TimeSheetRepository) GetStatusTimeSheetByID(id string) (*entity.StatusTimeSheet, error) {
	var status *entity.StatusTimeSheet
	err := config.DB.Where("id = ?", id).First(&status).Error
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (TimeSheetRepository) GetStatusTimeSheetByName(name string) (*entity.StatusTimeSheet, error) {
	var status *entity.StatusTimeSheet
	err := config.DB.Where("status_name = ?", name).First(&status).Error
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (TimeSheetRepository) GetTimeSheetByID(id string) (*entity.TimeSheet, error) {
	var ts entity.TimeSheet
	err := config.DB.Preload("TimeSheetDetails").First(&ts, "id = ?", id).Error
	return &ts, err
}

func (TimeSheetRepository) GetAllTimeSheets() (*[]entity.TimeSheet, error) {
	var timeSheets []entity.TimeSheet
	err := config.DB.Preload("TimeSheetDetails").Find(&timeSheets).Error
	return &timeSheets, err
}

func (TimeSheetRepository) ApproveManagerTimeSheet(id string, userID string) error {
	return config.DB.Model(&entity.TimeSheet{}).Where("id = ?", id).Update("confirmed_manager_by", userID).Error
}

func (TimeSheetRepository) RejectManagerTimeSheet(id string, userID string) error {
	return config.DB.Model(&entity.TimeSheet{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"updated_at":           time.Now(),
			"confirmed_manager_by": "",
		}).Error
}

func (TimeSheetRepository) ApproveBenefitTimeSheet(id string, userID string) error {
	return config.DB.Model(&entity.TimeSheet{}).Where("id = ?", id).Update("confirmed_benefit_by", userID).Error
}

func (TimeSheetRepository) RejectBenefitTimeSheet(id string, userID string) error {
	return config.DB.Model(&entity.TimeSheet{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"updated_at":           time.Now(),
			"confirmed_benefit_by": "",
		}).Error
}
