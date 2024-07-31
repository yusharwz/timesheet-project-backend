package service

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
	"final-project-enigma/repository"
)

type TimeSheetService interface {
	CreateTimeSheet(tsQ *request.TimeSheetRequest) (*response.TimeSheetResponse, error)
	FindByIdTimeSheet(id string) (*response.TimeSheetResponse, error)
	FindAllTimeSheet() (*[]response.TimeSheetResponse, error)
	UpdateTimeSheet(tsQ *request.TimeSheetRequest) (*response.TimeSheetResponse, error)
	DeleteTimeSheet(id string) error
}

type timeSheetService struct {
	tr repository.TimeSheetRepository
}

func (ts *timeSheetService) CreateTimeSheet(tsQ *request.TimeSheetRequest) (*response.TimeSheetResponse, error) {
	timeSheet := &entity.TimeSheet{
		Base: entity.Base{
			ID: tsQ.UserID,
		},
		ConfirmedManagerBy: tsQ.ConfirmedManagerBy,
		ConfirmedBenefitBy: tsQ.ConfirmedBenefitBy,
		StatusTimeSheetID:  tsQ.StatusTimeSheetID,
		UserID:             tsQ.UserID,
		TimeSheetDetails:   convertToTimeSheetDetails(tsQ.TimeSheetDetails),
	}

	createdTimeSheet, err := ts.tr.CreateTimeSheet(timeSheet)
	if err != nil {
		return nil, err
	}

	return convertToTimeSheetResponse(createdTimeSheet), nil
}
func (ts *timeSheetService) FindByIdTimeSheet(id string) (*response.TimeSheetResponse, error) {
	timeSheet, err := ts.tr.FindByIdTimeSheet(id)
	if err != nil {
		return nil, err
	}

	return convertToTimeSheetResponse(timeSheet), nil
}
func (ts *timeSheetService) FindAllTimeSheet() (*[]response.TimeSheetResponse, error) {
	timeSheets, err := ts.tr.FindAllTimeSheet()
	if err != nil {
		return nil, err
	}

	var timeSheetResponses []response.TimeSheetResponse
	for _, timeSheet := range *timeSheets {
		timeSheetResponses = append(timeSheetResponses, *convertToTimeSheetResponse(&timeSheet))
	}

	return &timeSheetResponses, nil
}
func (ts *timeSheetService) UpdateTimeSheet(tsQ *request.TimeSheetRequest) (*response.TimeSheetResponse, error) {
	timeSheet := &entity.TimeSheet{
		Base: entity.Base{
			ID: tsQ.UserID, // Assuming UserID is being used as an ID, usually it's better to generate a unique ID
		},
		ConfirmedManagerBy: tsQ.ConfirmedManagerBy,
		ConfirmedBenefitBy: tsQ.ConfirmedBenefitBy,
		StatusTimeSheetID:  tsQ.StatusTimeSheetID,
		UserID:             tsQ.UserID,
		TimeSheetDetails:   convertToTimeSheetDetails(tsQ.TimeSheetDetails),
	}

	updatedTimeSheet, err := ts.tr.UpdateTimeSheet(timeSheet)
	if err != nil {
		return nil, err
	}

	return convertToTimeSheetResponse(updatedTimeSheet), nil
}
func (ts *timeSheetService) DeleteTimeSheet(id string) error {
	return ts.tr.DeleteTimeSheet(id)
}

func NewTimeSheetServices(tr repository.TimeSheetRepository) TimeSheetService {
	return &timeSheetService{tr: tr}
}

// helper code
func convertToTimeSheetDetails(details []request.TimeSheetDetailRequest) []entity.TimeSheetDetail {
	var timeSheetDetails []entity.TimeSheetDetail
	for _, detail := range details {
		timeSheetDetails = append(timeSheetDetails, entity.TimeSheetDetail{
			Date:      detail.Date,
			StartTime: detail.StartTime,
			EndTime:   detail.EndTime,
			WorkID:    detail.WorkID,
		})
	}
	return timeSheetDetails
}

func convertToTimeSheetResponse(ts *entity.TimeSheet) *response.TimeSheetResponse {
	return &response.TimeSheetResponse{
		ID:                 ts.ID,
		ConfirmedManagerBy: ts.ConfirmedManagerBy,
		ConfirmedBenefitBy: ts.ConfirmedBenefitBy,
		StatusTimeSheetID:  ts.StatusTimeSheetID,
		UserID:             ts.UserID,
		TimeSheetDetails:   convertToTimeSheetDetailResponses(ts.TimeSheetDetails),
	}
}

func convertToTimeSheetDetailResponses(details []entity.TimeSheetDetail) []response.TimeSheetDetailResponse {
	var responses []response.TimeSheetDetailResponse
	for _, detail := range details {
		responses = append(responses, response.TimeSheetDetailResponse{
			ID:        detail.ID,
			Date:      detail.Date,
			StartTime: detail.StartTime,
			EndTime:   detail.EndTime,
			WorkID:    detail.WorkID,
		})
	}
	return responses
}
