package impl

import (
	"errors"
	"strconv"
	"timesheet-app/config"
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/entity"
	"timesheet-app/helper"
	"timesheet-app/repository"
	"timesheet-app/repository/impl"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type WorkService struct{}

var workRepository repository.WorkRepository = impl.NewWorkRepository()

func NewWorkService() *WorkService {
	return &WorkService{}
}

func (WorkService) CreateWork(request request.WorkRequest) (response.WorkResponse, error) {
	newWork := entity.Work{
		Base:             entity.Base{ID: uuid.NewString()},
		Description:      request.Description,
		Fee:              request.Fee,
		TimeSheetDetails: nil,
	}
	result, err := workRepository.CreateWork(newWork)
	if err != nil {
		log.Error().Msg(err.Error())
		return response.WorkResponse{}, err
	}
	return response.WorkResponse{
		Id:          result.ID,
		Description: result.Description,
		Fee:         result.Fee,
	}, nil
}

func (WorkService) UpdateWork(id string, request request.WorkRequest) (response.WorkResponse, error) {
	var existingWork entity.Work
	if err := config.DB.First(&existingWork, "id = ?", id).Error; err != nil {
		log.Error().Msg(err.Error())
		return response.WorkResponse{}, err
	}

	existingWork.Description = request.Description
	existingWork.Fee = request.Fee

	updatedWork, err := workRepository.UpdateWork(existingWork)
	if err != nil {
		log.Error().Msg(err.Error())
		return response.WorkResponse{}, err
	}

	return response.WorkResponse{
		Id:          updatedWork.ID,
		Description: updatedWork.Description,
		Fee:         updatedWork.Fee,
	}, nil
}

func (WorkService) DeleteWork(id string) error {
	if err := workRepository.DeleteWork(id); err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}

func (WorkService) GetById(id string, useSpec, getDeleted bool) (response.WorkResponse, error) {
	var spec func(db *gorm.DB) *gorm.DB
	if getDeleted {
		spec = func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}
	}
	result, err := workRepository.GetById(id, useSpec, spec)
	if err != nil {
		log.Error().Msg(err.Error())
		return response.WorkResponse{}, err
	}
	return response.WorkResponse{
		Id:          result.ID,
		Description: result.Description,
		Fee:         result.Fee,
	}, nil
}

func (WorkService) GetAllWork(paging, rowsPerPage, description string) ([]response.WorkResponse, string, string, error) {
	pagingInt, err := strconv.Atoi(paging)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, "0", "0", errors.New("invalid query for paging")
	}
	rowsPerPageInt, err := strconv.Atoi(rowsPerPage)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, "0", "0", errors.New("invalid query for rows per page")
	}

	var spec []func(db *gorm.DB) *gorm.DB
	spec = append(spec, helper.Paginate(pagingInt, rowsPerPageInt))
	if description != "" {
		spec = append(spec, helper.SelectWorkByDescription(description))
	}

	results, totalRows, err := workRepository.GetAllWork(spec)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, "0", "0", err
	}
	responses := make([]response.WorkResponse, 0)
	for _, v := range results {
		responses = append(responses, response.WorkResponse{
			Id:          v.ID,
			Description: v.Description,
			Fee:         v.Fee,
		})
	}
	totalPage := helper.GetTotalPage(totalRows, rowsPerPageInt)
	return responses, totalRows, strconv.Itoa(totalPage), nil
}
