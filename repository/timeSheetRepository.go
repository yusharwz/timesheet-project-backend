package repository

import (
	"final-project-enigma/entity"
	"log"
	"time"

	"gorm.io/gorm"
)

type TimeSheetRepository interface {
	CreateTimeSheet(ts *entity.TimeSheet) error
	UpdateTimeSheet(ts *entity.TimeSheet) error
	DeleteTimeSheet(id string) error
	GetTimeSheetByID(id string) (*entity.TimeSheet, error)
	GetAllTimeSheets() (*[]entity.TimeSheet, error)
	ApproveManagerTimeSheet(id string, userID string) error
	RejectManagerTimeSheet(id string, userID string) error
	ApproveBenefitTimeSheet(id string, userID string) error
	RejectBenefitTimeSheet(id string, userID string) error
}

type timeSheetRepository struct {
	db *gorm.DB
}

func NewTimeSheetRepository(db *gorm.DB) TimeSheetRepository {
	return &timeSheetRepository{db: db}
}

func (r *timeSheetRepository) CreateTimeSheet(ts *entity.TimeSheet) error {
	return r.db.Create(ts).Error
}

func (r *timeSheetRepository) UpdateTimeSheet(ts *entity.TimeSheet) error {
	return r.db.Save(ts).Error
}

func (r *timeSheetRepository) DeleteTimeSheet(id string) error {
	log.Printf("Deleting time sheet with ID: %s", id)
	return r.db.Model(&entity.TimeSheet{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
}

func (r *timeSheetRepository) GetTimeSheetByID(id string) (*entity.TimeSheet, error) {
	var ts entity.TimeSheet
	err := r.db.Preload("TimeSheetDetails").First(&ts, "id = ?", id).Error
	return &ts, err
}

func (r *timeSheetRepository) GetAllTimeSheets() (*[]entity.TimeSheet, error) {
	var timeSheets []entity.TimeSheet
	err := r.db.Preload("TimeSheetDetails").Find(&timeSheets).Error
	return &timeSheets, err
}

func (r *timeSheetRepository) ApproveManagerTimeSheet(id string, userID string) error {
	return r.db.Model(&entity.TimeSheet{}).Where("id = ?", id).Update("confirmed_manager_by", userID).Error
}

func (r *timeSheetRepository) RejectManagerTimeSheet(id string, userID string) error {
	return r.db.Model(&entity.TimeSheet{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"updated_at":           time.Now(),
			"confirmed_manager_by": "",
		}).Error
}

func (r *timeSheetRepository) ApproveBenefitTimeSheet(id string, userID string) error {
	return r.db.Model(&entity.TimeSheet{}).Where("id = ?", id).Update("confirmed_benefit_by", userID).Error
}

func (r *timeSheetRepository) RejectBenefitTimeSheet(id string, userID string) error {
	return r.db.Model(&entity.TimeSheet{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"updated_at":           time.Now(),
			"confirmed_benefit_by": "",
		}).Error
}

