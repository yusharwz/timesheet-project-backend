package impl

import (
	"errors"
	"final-project-enigma/config"
	"final-project-enigma/dto/request"
	"final-project-enigma/entity"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type TimeSheetRepository struct{}

func NewTimeSheetRepository() *TimeSheetRepository {
	return &TimeSheetRepository{}
}

func (TimeSheetRepository) CreateTimeSheet(timesheet entity.TimeSheet) (*entity.TimeSheet, error) {
	err := config.DB.Create(&timesheet).Error
	if err != nil {
		return nil, err
	}
	return &timesheet, nil
}

func (TimeSheetRepository) UpdateTimeSheet(ts entity.TimeSheet) (*entity.TimeSheet, error) {
	err := config.DB.Save(&ts).Error
	if err != nil {
		return nil, err
	}
	return &ts, nil
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

func (TimeSheetRepository) GetAllTimeSheets(spec []func(db *gorm.DB) *gorm.DB) (*[]entity.TimeSheet, string, error) {
	var timeSheets []entity.TimeSheet
	err := config.DB.Scopes(spec...).Preload("TimeSheetDetails").Find(&timeSheets).Error
	if len(timeSheets) == 0 {
		return nil, "0", errors.New("data not found")
	}
	var totalRows int64
	config.DB.Model(&entity.TimeSheet{}).Count(&totalRows)
	return &timeSheets, strconv.FormatInt(totalRows, 10), err
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

func (TimeSheetRepository) UpdateTimeSheetStatus(req request.UpdateTimeSheetStatusRequest) error {
	var timeSheet entity.TimeSheet
	var timeSheetsStatus entity.StatusTimeSheet

	if err := config.DB.Where("status_name = ?", "pending").First(&timeSheetsStatus).Error; err != nil {
		return err
	}

	if err := config.DB.Model(&timeSheet).Where("id = ?", req.TimeSheetID).Update("status_timesheet_id", timeSheetsStatus.ID).Error; err != nil {
		return err
	}

	return nil
}
