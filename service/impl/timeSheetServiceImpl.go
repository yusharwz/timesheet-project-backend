package impl

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/entity"
	"timesheet-app/helper"
	"timesheet-app/repository"
	"timesheet-app/repository/impl"
	"timesheet-app/service"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
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
		log.Error().Msg(err.Error())
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
		log.Error().Msg(err.Error())
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
		log.Error().Msg(err.Error())
		return nil, err
	}

	timeSheetDetailsResponse := make([]response.TimeSheetDetailResponse, 0)
	var total int
	for _, v := range timeSheetDetails {
		var fee int
		work, err := workService.GetById(v.WorkID, false)
		if err != nil {
			log.Error().Msg(err.Error())
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
		Status:             "created",
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
		log.Error().Msg(err.Error())
		return nil, err
	}

	status, err := timeSheetRepository.GetStatusTimeSheetByName("created")
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	if existingTs.StatusTimeSheetID != status.ID {
		return nil, errors.New("timesheet cannot be updated as it has been approve by manager")
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
		log.Error().Msg(err.Error())
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
		log.Error().Msg(err.Error())
		return nil, err
	}
	timeSheetDetailsResponse := make([]response.TimeSheetDetailResponse, 0)
	var total int
	for _, v := range res.TimeSheetDetails {
		var fee int
		work, err := workService.GetById(v.WorkID, false)
		if err != nil {
			log.Error().Msg(err.Error())
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

	statusName, err := timeSheetRepository.GetStatusTimeSheetByID(existingTs.StatusTimeSheetID)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	timeSheetResponse := response.TimeSheetResponse{
		ID:                 res.ID,
		CreatedAt:          res.CreatedAt,
		UpdatedAt:          res.UpdatedAt,
		Status:             statusName.StatusName,
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
		log.Error().Msg(err.Error())
		return err
	}

	status, err := timeSheetRepository.GetStatusTimeSheetByName("created")
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	if existingTs.StatusTimeSheetID != status.ID {
		return errors.New("timesheet cannot be deleted as it has been approved or rejected")
	}

	return timeSheetRepository.DeleteTimeSheet(id)
}

func (TimeSheetService) GetTimeSheetByID(id string) (*response.TimeSheetResponse, error) {
	res, err := timeSheetRepository.GetTimeSheetByID(id)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	timeSheetDetailsResponse := make([]response.TimeSheetDetailResponse, 0)
	var total int
	for _, v := range res.TimeSheetDetails {
		var fee int
		work, err := workService.GetById(v.WorkID, true)
		if err != nil {
			log.Error().Msg(err.Error())
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
		log.Error().Msg(err.Error())
		return nil, err
	}

	status, err := timeSheetRepository.GetStatusTimeSheetByID(res.StatusTimeSheetID)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	var managerResponse response.ConfirmedByResponse
	var benefitResponse response.ConfirmedByResponse
	if res.ConfirmedManagerBy != "" {
		manager, err := accountService.GetAccountByID(res.ConfirmedManagerBy)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, err
		}
		managerResponse.ID = manager.UserID
		managerResponse.Name = manager.Name
		managerResponse.Email = manager.Email
		managerResponse.SignatureUrl = manager.SignatureURL
	}

	if res.ConfirmedBenefitBy != "" {
		benefit, err := accountService.GetAccountByID(res.ConfirmedBenefitBy)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, err
		}
		benefitResponse.ID = benefit.UserID
		benefitResponse.Name = benefit.Name
		benefitResponse.Email = benefit.Email
		benefitResponse.SignatureUrl = benefit.SignatureURL
	}

	timeSheetResponse := response.TimeSheetResponse{
		ID:                 res.ID,
		CreatedAt:          res.CreatedAt,
		UpdatedAt:          res.UpdatedAt,
		Status:             status.StatusName,
		ConfirmedManagerBy: managerResponse,
		ConfirmedBenefitBy: benefitResponse,
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

func (TimeSheetService) GetAllTimeSheets(paging, rowsPerPage, year, userId, name string, period, status []string) (*[]response.TimeSheetResponse, string, string, error) {
	var err error
	var pagingInt int
	var rowsPerPageInt int
	var totalRows string
	var spec []func(db *gorm.DB) *gorm.DB
	var results *[]entity.TimeSheet
	listStatusId := []string{"", "", ""}

	pagingInt, err = strconv.Atoi(paging)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, "0", "0", errors.New("invalid query for paging")
	}
	rowsPerPageInt, err = strconv.Atoi(rowsPerPage)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, "0", "0", errors.New("invalid query for rows per page")
	}

	spec = append(spec, helper.Paginate(pagingInt, rowsPerPageInt))
	if year != "" && period != nil {
		spec = append(spec, helper.SelectByPeriod(year, period[0], period[1]))
	}

	if userId != "" {
		spec = append(spec, helper.SelectByUserId(userId))
	}

	if status != nil {
		var res *entity.StatusTimeSheet
		var err error
		for i := 0; i < len(status); i++ {
			res, err = timeSheetRepository.GetStatusTimeSheetByName(status[i])
			if err != nil {
				return nil, "0", "0", err
			}
			listStatusId[i] = res.ID
		}
		spec = append(spec, helper.SelectByStatus(listStatusId))
	}

	if name != "" {
		spec = append(spec, helper.SelectUserInTimeSheet(name))
	}

	results, totalRows, err = timeSheetRepository.GetAllTimeSheets(spec)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, "0", "0", err
	}
	timeSheetsResponse := make([]response.TimeSheetResponse, 0)

	for _, v := range *results {
		status, err := timeSheetRepository.GetStatusTimeSheetByID(v.StatusTimeSheetID)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, "0", "0", err
		}
		user, err := accountService.GetAccountByID(v.UserID)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, "0", "0", err
		}

		timeSheetDetailsResponse := make([]response.TimeSheetDetailResponse, 0)
		var total int
		for _, v := range v.TimeSheetDetails {
			var fee int
			work, err := workService.GetById(v.WorkID, true)
			if err != nil {
				log.Error().Msg(err.Error())
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

		var managerResponse response.ConfirmedByResponse
		var benefitResponse response.ConfirmedByResponse
		if v.ConfirmedManagerBy != "" {
			manager, err := accountService.GetAccountByID(v.ConfirmedManagerBy)
			if err != nil {
				log.Error().Msg(err.Error())
				return nil, "0", "0", err
			}
			managerResponse.ID = manager.UserID
			managerResponse.Name = manager.Name
			managerResponse.Email = manager.Email
			managerResponse.SignatureUrl = manager.SignatureURL
		}

		if v.ConfirmedBenefitBy != "" {
			benefit, err := accountService.GetAccountByID(v.ConfirmedBenefitBy)
			if err != nil {
				log.Error().Msg(err.Error())
				return nil, "0", "0", err
			}
			benefitResponse.ID = benefit.UserID
			benefitResponse.Name = benefit.Name
			benefitResponse.Email = benefit.Email
			benefitResponse.SignatureUrl = benefit.SignatureURL
		}

		timeSheetsResponse = append(timeSheetsResponse, response.TimeSheetResponse{
			ID:                 v.ID,
			CreatedAt:          v.CreatedAt,
			UpdatedAt:          v.UpdatedAt,
			Status:             status.StatusName,
			ConfirmedManagerBy: managerResponse,
			ConfirmedBenefitBy: benefitResponse,
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

	timeSheet, err := timeSheetRepository.GetTimeSheetByID(id)
	if err != nil {
		log.Error().Msgf("Failed to get timesheet by ID: %v", err)
		return err
	}

	status, err := timeSheetRepository.GetStatusTimeSheetByName("pending")
	if err != nil {
		log.Error().Msgf("Failed to get status 'pending': %v", err)
		return err
	}

	if timeSheet.StatusTimeSheetID == status.ID {
		err = timeSheetRepository.ApproveManagerTimeSheet(id, userID)
		if err != nil {
			log.Error().Msgf("Failed to approve timesheet: %v", err)
			return err
		}
	} else if timeSheet.ConfirmedManagerBy != "" {
		return errors.New("timesheet has already been submitted")
	} else {
		return errors.New("timesheet not submitted")
	}

	detailAccount, detailUser, err := accountRepository.GetAccountDetailByUserID(timeSheet.UserID)
	if err != nil {
		log.Error().Msgf("Failed to get account details by userID: %v", err)
		return err
	}

	err = helper.SendNotificationToTrainer(detailAccount.Email, detailUser.Name, "Accepted", "Manager")
	if err != nil {
		log.Error().Msgf("Failed to send approval email: %v", err)
		return err
	}

	emails, err := timeSheetRepository.GetBenefitEmails()
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	for _, email := range emails {
		err := helper.SendNotificationToBenefit(email, detailUser.Name)
		if err != nil {
			log.Error().Msg(err.Error())
			continue
		}
	}

	return nil
}

func (TimeSheetService) RejectManagerTimeSheet(id string, userID string) error {

	timeSheet, err := timeSheetRepository.GetTimeSheetByID(id)
	if err != nil {
		log.Error().Msgf("Failed to get timesheet by ID: %v", err)
		return err
	}

	status, err := timeSheetRepository.GetStatusTimeSheetByName("pending")
	if err != nil {
		log.Error().Msgf("Failed to get status 'pending': %v", err)
		return err
	}

	if timeSheet.ConfirmedManagerBy != "" {
		return errors.New("timesheet has already been submitted")
	}

	if timeSheet.StatusTimeSheetID == status.ID {
		err = timeSheetRepository.RejectManagerTimeSheet(id, userID)
		if err != nil {
			log.Error().Msgf("Failed to reject timesheet: %v", err)
			return err
		}
	}

	detailAccount, detailUser, err := accountRepository.GetAccountDetailByUserID(timeSheet.UserID)
	if err != nil {
		log.Error().Msgf("Failed to get account details by userID: %v", err)
		return err
	}

	err = helper.SendNotificationToTrainer(detailAccount.Email, detailUser.Name, "Denied", "Manager")
	if err != nil {
		log.Error().Msgf("Failed to send rejection email: %v", err)
		return err
	}

	return nil
}

func (TimeSheetService) ApproveBenefitTimeSheet(id string, userID string) error {
	timeSheet, err := timeSheetRepository.GetTimeSheetByID(id)
	if err != nil {
		log.Error().Msgf("Failed to get timesheet by ID: %v", err)
		return err
	}

	status, err := timeSheetRepository.GetStatusTimeSheetByName("accepted")
	if err != nil {
		log.Error().Msgf("Failed to get status 'pending': %v", err)
		return err
	}

	if timeSheet.StatusTimeSheetID == status.ID {
		err = timeSheetRepository.ApproveBenefitTimeSheet(id, userID)
		if err != nil {
			log.Error().Msgf("Failed to approve timesheet: %v", err)
			return err
		}
	} else if timeSheet.ConfirmedBenefitBy != "" {
		return errors.New("timesheet has already been submitted")
	} else {
		return errors.New("timesheet not submitted")
	}

	detailAccount, detailUser, err := accountRepository.GetAccountDetailByUserID(timeSheet.UserID)
	if err != nil {
		log.Error().Msgf("Failed to get account details by userID: %v", err)
		return err
	}

	err = helper.SendNotificationToTrainer(detailAccount.Email, detailUser.Name, "Approved", "Benefit")
	if err != nil {
		log.Error().Msgf("Failed to send approval email: %v", err)
		return err
	}

	return nil
}

func (TimeSheetService) RejectBenefitTimeSheet(id string, userID string) error {
	timeSheet, err := timeSheetRepository.GetTimeSheetByID(id)
	if err != nil {
		log.Error().Msgf("Failed to get timesheet by ID: %v", err)
		return err
	}

	status, err := timeSheetRepository.GetStatusTimeSheetByName("accepted")
	if err != nil {
		log.Error().Msgf("Failed to get status 'pending': %v", err)
		return err
	}

	if timeSheet.StatusTimeSheetID == status.ID {
		err = timeSheetRepository.RejectBenefitTimeSheet(id, userID)
		if err != nil {
			log.Error().Msgf("Failed to approve timesheet: %v", err)
			return err
		}
	} else if timeSheet.ConfirmedBenefitBy != "" {
		return errors.New("timesheet has already been submitted")
	} else {
		return errors.New("timesheet not submitted")
	}

	detailAccount, detailUser, err := accountRepository.GetAccountDetailByUserID(timeSheet.UserID)
	if err != nil {
		log.Error().Msgf("Failed to get account details by userID: %v", err)
		return err
	}

	err = helper.SendNotificationToTrainer(detailAccount.Email, detailUser.Name, "Rejected", "Benefit")
	if err != nil {
		log.Error().Msgf("Failed to send approval email: %v", err)
		return err
	}

	return nil
}

func (TimeSheetService) UpdateTimeSheetStatus(id string) error {
	timeNow := time.Now()
	day := timeNow.Day()

	if day != 19 && day != 20 {
		log.Error().Msg("failed to update time sheet status, please only submit on 19 or 20")
		return errors.New("failed to update time sheet status, please only submit on 19 or 20")
	}

	err := timeSheetRepository.UpdateTimeSheetStatus(id)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	timeSheet, err := timeSheetRepository.GetTimeSheetByID(id)
	if err != nil {
		log.Error().Msgf("Failed to get timesheet by ID: %v", err)
		return err
	}

	_, detailUser, err := accountRepository.GetAccountDetailByUserID(timeSheet.UserID)
	if err != nil {
		log.Error().Msgf("Failed to get account details by userID: %v", err)
		return err
	}

	emails, err := timeSheetRepository.GetManagerEmails()
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	for _, email := range emails {
		err := helper.SendNotificationToManager(email, detailUser.Name)
		if err != nil {
			log.Error().Msg(err.Error())
			continue
		}
	}

	return nil
}
