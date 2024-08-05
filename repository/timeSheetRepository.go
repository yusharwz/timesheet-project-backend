package repository

import (
	"final-project-enigma/entity"
)

type TimeSheetRepository interface {
	CreateTimeSheet(ts entity.TimeSheet) (*entity.TimeSheet, error)
	UpdateTimeSheet(ts entity.TimeSheet) (*entity.TimeSheet, error)
	DeleteTimeSheet(id string) error
	GetStatusTimeSheetByID(id string) (*entity.StatusTimeSheet, error)
	GetStatusTimeSheetByName(name string) (*entity.StatusTimeSheet, error)
	GetTimeSheetByID(id string) (*entity.TimeSheet, error)
	GetAllTimeSheets(paging, rowsPerPage int) (*[]entity.TimeSheet, error)
	ApproveManagerTimeSheet(id string, userID string) error
	RejectManagerTimeSheet(id string, userID string) error
	ApproveBenefitTimeSheet(id string, userID string) error
	RejectBenefitTimeSheet(id string, userID string) error
}
