package impl

import (
	"errors"
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
	"final-project-enigma/helper"
	"final-project-enigma/repository"
	"final-project-enigma/repository/impl"
	"final-project-enigma/service"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type TimeSheetService struct{}

var timeSheetRepository repository.TimeSheetRepository = impl.NewTimeSheetRepository()
var accountService service.AccountService = NewAccountService()
var workService service.WorkService = NewWorkService()

func NewTimeSheetService() *TimeSheetService {
	return &TimeSheetService{}
}

func (TimeSheetService) CreateTimeSheet(req request.TimeSheetRequest, authHeader string) (*response.TimeSheetResponse, error) {
	status, err := timeSheetRepository.GetStatusTimeSheetByName("created")
	if err != nil {
		return nil, err
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

	user, err := accountService.GetAccountDetail(authHeader)
	if err != nil {
		return nil, err
	}

	timeSheet := entity.TimeSheet{
		Base:              entity.Base{ID: uuid.NewString()},
		StatusTimeSheetID: status.ID,
		UserID:            user.UserID,
		TimeSheetDetails:  timeSheetDetails,
	}

	res, err := timeSheetRepository.CreateTimeSheet(timeSheet)
	if err != nil {
		return nil, err
	}

	timeSheetDetailsResponse := make([]response.TimeSheetDetailResponse, 0)
	var total int
	for _, v := range timeSheetDetails {
		var fee int
		work, err := workService.GetById(v.WorkID)
		if err != nil {
			return nil, err
		}
		duration := int(v.EndTime.Sub(v.StartTime).Hours())
		if duration < 1 {
			return nil, errors.New("invalid work duration")
		}
		if strings.Contains(strings.ToLower(work.Description), "interview") && duration >= 2 {
			fee = 50000
		} else {
			fee = work.Fee
		}
		subTotal := fee * duration
		total += subTotal
		timeSheetDetailsResponse = append(timeSheetDetailsResponse, response.TimeSheetDetailResponse{
			ID:          v.ID,
			Date:        v.Date,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			WorkID:      v.WorkID,
			Description: work.Description,
			SubTotal:    subTotal,
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
		UserTimeSheetResponse: response.UserTimeSheetResponse{
			ID:           user.UserID,
			Name:         user.Name,
			Email:        user.Email,
			SignatureUrl: user.SignatureURL,
		},
		TimeSheetDetails: timeSheetDetailsResponse,
		Total:            total,
	}

	return &timeSheetResponse, nil
}

func (TimeSheetService) UpdateTimeSheet(req request.UpdateTimeSheetRequest, authHeader string) (*response.TimeSheetResponse, error) {
	existingTs, err := timeSheetRepository.GetTimeSheetByID(req.ID)
	if err != nil {
		return nil, err
	}

	status, err := timeSheetRepository.GetStatusTimeSheetByName("created")
	if err != nil {
		return nil, err
	}

	if existingTs.StatusTimeSheetID != status.ID {
		return nil, errors.New("timesheet cannot be updated as it has been approved or rejected")
	}

	timeSheetDetails := make([]entity.TimeSheetDetail, 0)
	for _, value := range req.TimeSheetDetails {
		timeSheetDetails = append(timeSheetDetails, entity.TimeSheetDetail{
			Base:      entity.Base{ID: value.ID},
			Date:      value.Date,
			StartTime: value.StartTime,
			EndTime:   value.EndTime,
			WorkID:    value.WorkID,
		})
	}

	user, err := accountService.GetAccountDetail(authHeader)
	if err != nil {
		return nil, err
	}

	timeSheet := entity.TimeSheet{
		Base:              entity.Base{ID: req.ID},
		StatusTimeSheetID: status.ID,
		UserID:            user.UserID,
		TimeSheetDetails:  timeSheetDetails,
	}

	res, err := timeSheetRepository.UpdateTimeSheet(timeSheet)
	if err != nil {
		return nil, err
	}

	timeSheetDetailsResponse := make([]response.TimeSheetDetailResponse, 0)
	var total int
	for _, v := range timeSheetDetails {
		var fee int
		work, err := workService.GetById(v.WorkID)
		if err != nil {
			return nil, err
		}
		duration := int(v.EndTime.Sub(v.StartTime).Hours())
		if duration < 1 {
			return nil, errors.New("invalid work duration")
		}
		if strings.Contains(strings.ToLower(work.Description), "interview") && duration >= 2 {
			fee = 50000
		} else {
			fee = work.Fee
		}
		subTotal := fee * duration
		total += subTotal
		timeSheetDetailsResponse = append(timeSheetDetailsResponse, response.TimeSheetDetailResponse{
			ID:          v.ID,
			Date:        v.Date,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			WorkID:      v.WorkID,
			Description: work.Description,
			SubTotal:    subTotal,
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
		UserTimeSheetResponse: response.UserTimeSheetResponse{
			ID:           user.UserID,
			Name:         user.Name,
			Email:        user.Email,
			SignatureUrl: user.SignatureURL,
		},
		TimeSheetDetails: timeSheetDetailsResponse,
		Total:            total,
	}
	return &timeSheetResponse, nil
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

func (TimeSheetService) GetTimeSheetByID(id string) (*response.TimeSheetResponse, error) {
	res, err := timeSheetRepository.GetTimeSheetByID(id)
	if err != nil {
		return nil, err
	}
	timeSheetDetailsResponse := make([]response.TimeSheetDetailResponse, 0)
	var total int
	for _, v := range res.TimeSheetDetails {
		var fee int
		work, err := workService.GetById(v.WorkID)
		if err != nil {
			return nil, err
		}
		duration := int(v.EndTime.Sub(v.StartTime).Hours())
		if duration < 1 {
			return nil, errors.New("invalid work duration")
		}
		if strings.Contains(strings.ToLower(work.Description), "interview") && duration >= 2 {
			fee = 50000
		} else {
			fee = work.Fee
		}
		subTotal := fee * duration
		total += subTotal
		timeSheetDetailsResponse = append(timeSheetDetailsResponse, response.TimeSheetDetailResponse{
			ID:          v.ID,
			Date:        v.Date,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			WorkID:      v.WorkID,
			Description: work.Description,
			SubTotal:    subTotal,
		})
	}

	user, err := accountService.GetAccountByID(res.UserID)
	if err != nil {
		return nil, err
	}

	timeSheetResponse := response.TimeSheetResponse{
		ID:                 res.ID,
		CreatedAt:          res.CreatedAt,
		UpdatedAt:          res.UpdatedAt,
		StatusByManager:    "created",
		StatusByBenefit:    "created",
		ConfirmedManagerBy: response.ConfirmedByResponse{},
		ConfirmedBenefitBy: response.ConfirmedByResponse{},
		UserTimeSheetResponse: response.UserTimeSheetResponse{
			ID:           user.UserID,
			Name:         user.Name,
			Email:        user.Email,
			SignatureUrl: user.SignatureURL,
		},
		TimeSheetDetails: timeSheetDetailsResponse,
		Total:            total,
	}
	return &timeSheetResponse, nil
}

func (TimeSheetService) GetAllTimeSheets(paging, rowsPerPage, period, userId, confirm, status string) (*[]response.TimeSheetResponse, string, string, error) {
	var err error
	var pagingInt int
	var rowsPerPageInt int
	var totalRows string
	var results *[]entity.TimeSheet

	pagingInt, err = strconv.Atoi(paging)
	if err != nil {
		return nil, "0", "0", errors.New("invalid query for paging")
	}
	rowsPerPageInt, err = strconv.Atoi(rowsPerPage)
	if err != nil {
		return nil, "0", "0", errors.New("invalid query for rows per page")
	}

	results, totalRows, err = timeSheetRepository.GetAllTimeSheets(pagingInt, rowsPerPageInt)
	if err != nil {
		return nil, "0", "0", err
	}
	timeSheetsResponse := make([]response.TimeSheetResponse, 0)

	for _, v := range *results {
		var statusByManager string
		var statusByBenefit string

		status, err := timeSheetRepository.GetStatusTimeSheetByID(v.StatusTimeSheetID)
		if err != nil {
			return nil, "0", "0", err
		} else if (status.StatusName == "approved" || status.StatusName == "rejected") && v.ConfirmedManagerBy != "" {
			statusByManager = v.ConfirmedManagerBy
			statusByBenefit = "pending"
		}

		if v.ConfirmedManagerBy != "" && v.ConfirmedBenefitBy == "" {
			statusByManager = status.StatusName
			statusByBenefit = "pending"
			if v.ConfirmedBenefitBy != "" {
				statusByBenefit = status.StatusName
			}
		}

		user, err := accountService.GetAccountByID(v.UserID)
		if err != nil {
			return nil, "0", "0", err
		}

		timeSheetDetailsResponse := make([]response.TimeSheetDetailResponse, 0)
		var total int
		for _, v := range v.TimeSheetDetails {
			var fee int
			work, err := workService.GetById(v.WorkID)
			if err != nil {
				return nil, "0", "0", err
			}
			duration := int(v.EndTime.Sub(v.StartTime).Hours())
			if duration < 1 {
				return nil, "0", "0", errors.New("invalid work duration")
			}
			if strings.Contains(strings.ToLower(work.Description), "interview") && duration >= 2 {
				fee = 50000
			} else {
				fee = work.Fee
			}
			subTotal := fee * duration
			total += subTotal
			timeSheetDetailsResponse = append(timeSheetDetailsResponse, response.TimeSheetDetailResponse{
				ID:          v.ID,
				Date:        v.Date,
				StartTime:   v.StartTime,
				EndTime:     v.EndTime,
				WorkID:      v.WorkID,
				Description: work.Description,
				SubTotal:    subTotal,
			})
		}

		timeSheetsResponse = append(timeSheetsResponse, response.TimeSheetResponse{
			ID:                 v.ID,
			CreatedAt:          v.CreatedAt,
			UpdatedAt:          v.UpdatedAt,
			StatusByManager:    statusByManager,
			StatusByBenefit:    statusByBenefit,
			ConfirmedManagerBy: response.ConfirmedByResponse{},
			ConfirmedBenefitBy: response.ConfirmedByResponse{},
			UserTimeSheetResponse: response.UserTimeSheetResponse{
				ID:           user.UserID,
				Name:         user.Name,
				Email:        user.Email,
				SignatureUrl: user.SignatureURL,
			},
			TimeSheetDetails: timeSheetDetailsResponse,
			Total:            total,
		})
	}

	totalPage := helper.GetTotalPage(totalRows, rowsPerPageInt)
	return &timeSheetsResponse, totalRows, strconv.Itoa(totalPage), nil
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
