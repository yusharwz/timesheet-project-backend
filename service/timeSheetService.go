package service

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
)

type TimeSheetService interface {
	CreateTimeSheet(ts request.TimeSheetRequest, authHeader string) (*response.TimeSheetResponse, error)
	UpdateTimeSheet(ts request.UpdateTimeSheetRequest, authHeader string) (*response.TimeSheetResponse, error)
	DeleteTimeSheet(id string) error
	GetTimeSheetByID(id string) (*response.TimeSheetResponse, error)
	GetAllTimeSheets(paging, rowsPerPage, year, userId, status string, period []string) (*[]response.TimeSheetResponse, string, string, error)
	ApproveManagerTimeSheet(id string, userID string) error
	RejectManagerTimeSheet(id string, userID string) error
	ApproveBenefitTimeSheet(id string, userID string) error
	RejectBenefitTimeSheet(id string, userID string) error
	UpdateTimeSheetStatus(req request.UpdateTimeSheetStatusRequest) error
}
