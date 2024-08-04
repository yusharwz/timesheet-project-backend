package service

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
)

type TimeSheetService interface {
	CreateTimeSheet(ts request.TimeSheetRequest, authHeader string) (response.TimeSheetResponse, error)
	UpdateTimeSheet(ts *entity.TimeSheet) error
	DeleteTimeSheet(id string) error
	GetTimeSheetByID(id string) (*entity.TimeSheet, error)
	GetAllTimeSheets() (*[]entity.TimeSheet, error)
	ApproveManagerTimeSheet(id string, userID string) error
	RejectManagerTimeSheet(id string, userID string) error
	ApproveBenefitTimeSheet(id string, userID string) error
	RejectBenefitTimeSheet(id string, userID string) error
}
