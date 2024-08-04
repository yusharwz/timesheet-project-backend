package impl

import (
	"errors"
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
	"final-project-enigma/middleware"
	"final-project-enigma/repository"
	"final-project-enigma/repository/impl"
	"github.com/google/uuid"
)

type TimeSheetService struct{}

var timeSheetRepository repository.TimeSheetRepository = impl.NewTimeSheetRepository()

func NewTimeSheetService() *TimeSheetService {
	return &TimeSheetService{}
}

func (TimeSheetService) CreateTimeSheet(req request.TimeSheetRequest, authHeader string) (response.TimeSheetResponse, error) {
	status, err := timeSheetRepository.GetStatusTimeSheetByName("created")
	if err != nil {
		return response.TimeSheetResponse{}, err
	}

	timeSheetDetails := make([]entity.TimeSheetDetail, 0)
	for _, value := range req.TimeSheetDetails {
		timeSheetDetails = append(timeSheetDetails, entity.TimeSheetDetail{
			Base:      entity.Base{ID: uuid.NewString()},
			Date:      value.Date,
			StartTime: value.StartTime,
			EndTime:   value.EndTime,
			WorkID:    value.WorkID,
		})
	}

	userID, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		return response.TimeSheetResponse{}, err
	}

	timeSheet := entity.TimeSheet{
		Base:              entity.Base{ID: uuid.NewString()},
		StatusTimeSheetID: status.ID,
		UserID:            userID,
		TimeSheetDetails:  timeSheetDetails,
	}

	res, err := timeSheetRepository.CreateTimeSheet(timeSheet)
	if err != nil {
		return response.TimeSheetResponse{}, err
	}

	timeSheetDetailsResponse := make([]response.TimeSheetDetailResponse, 0)
	for _, v := range timeSheetDetails {
		timeSheetDetailsResponse = append(timeSheetDetailsResponse, response.TimeSheetDetailResponse{
			ID:        v.ID,
			Date:      v.Date,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
			WorkID:    v.WorkID,
		})
	}

	timeSheetResponse := response.TimeSheetResponse{
		ID:                 res.ID,
		CreatedAt:          res.CreatedAt,
		UpdatedAt:          res.UpdatedAt,
		StatusByManager:    "created",
		StatusByBenefit:    "created",
		ConfirmedManagerBy: response.ConfirmedByResponse{},
		ConfirmedBenefitBy: response.ConfirmedByResponse{},
		TimeSheetDetails:   timeSheetDetailsResponse,
	}

	return timeSheetResponse, nil
}

func (TimeSheetService) UpdateTimeSheet(ts *entity.TimeSheet) error {
	existingTs, err := timeSheetRepository.GetTimeSheetByID(ts.ID)
	if err != nil {
		return err
	}

	if existingTs.ConfirmedManagerBy != "" || existingTs.ConfirmedBenefitBy != "" {
		return errors.New("timesheet cannot be updated as it has been approved or rejected")
	}

	return timeSheetRepository.UpdateTimeSheet(ts)
}

func (TimeSheetService) DeleteTimeSheet(id string) error {
	existingTs, err := timeSheetRepository.GetTimeSheetByID(id)
	if err != nil {
		return err
	}

	if existingTs.ConfirmedManagerBy != "" || existingTs.ConfirmedBenefitBy != "" {
		return errors.New("timesheet cannot be deleted as it has been approved or rejected")
	}

	return timeSheetRepository.DeleteTimeSheet(id)
}

func (TimeSheetService) GetTimeSheetByID(id string) (*entity.TimeSheet, error) {
	return timeSheetRepository.GetTimeSheetByID(id)
}

func (TimeSheetService) GetAllTimeSheets() (*[]entity.TimeSheet, error) {
	return timeSheetRepository.GetAllTimeSheets()
}

func (TimeSheetService) ApproveManagerTimeSheet(id string, userID string) error {
	return timeSheetRepository.ApproveManagerTimeSheet(id, userID)
}

func (TimeSheetService) RejectManagerTimeSheet(id string, userID string) error {
	return timeSheetRepository.RejectManagerTimeSheet(id, userID)
}

func (TimeSheetService) ApproveBenefitTimeSheet(id string, userID string) error {
	return timeSheetRepository.ApproveBenefitTimeSheet(id, userID)
}

func (TimeSheetService) RejectBenefitTimeSheet(id string, userID string) error {
	return timeSheetRepository.RejectBenefitTimeSheet(id, userID)
}
