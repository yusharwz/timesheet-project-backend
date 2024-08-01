package impl

import (
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

func (WorkService) GetAllWork(page, size string) ([]response.WorkResponse, string, error) {
	results, total, err := workRepository.GetAllWork(page, size)
	if err != nil {
		return []response.WorkResponse{}, "0", err
	}
	responses := make([]response.WorkResponse, 0)
	for _, v := range results {
		responses = append(responses, response.WorkResponse{
			Id:          v.ID,
			Description: v.Description,
			Fee:         v.Fee,
		})
	}
	return responses, total, nil
}
