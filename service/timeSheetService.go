package service

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
)

type TimeSheetService interface {
	CreateTimeSheet(ts request.TimeSheetRequest, authHeader string) (*response.TimeSheetResponse, error)
	UpdateTimeSheet(ts request.UpdateTimeSheetRequest, authHeader string) (*response.TimeSheetResponse, error)
	DeleteTimeSheet(id string) error
	GetTimeSheetByID(id string) (*response.TimeSheetResponse, error)
	GetAllTimeSheets() (*[]entity.TimeSheet, error)
	ApproveManagerTimeSheet(id string, userID string) error
	RejectManagerTimeSheet(id string, userID string) error
	ApproveBenefitTimeSheet(id string, userID string) error
	RejectBenefitTimeSheet(id string, userID string) error
}
