package impl

import (
	"final-project-enigma/config"
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
	"final-project-enigma/repository/impl"

	"github.com/google/uuid"
)

type WorkService struct{}

var workRepository = impl.NewWorkRepository()

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
		return response.WorkResponse{}, err
	}

	existingWork.Description = request.Description
	existingWork.Fee = request.Fee

	updatedWork, err := workRepository.UpdateWork(existingWork)
	if err != nil {
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
		return err
	}
	return nil
}

func (WorkService) GetById(id string) (response.WorkResponse, error) {
	result, err := workRepository.GetById(id)
	if err != nil {
		return response.WorkResponse{}, err
	}
	return response.WorkResponse{
		Id:          result.ID,
		Description: result.Description,
		Fee:         result.Fee,
	}, nil
}
