package service

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
)

type WorkService interface {
	CreateWork(request request.WorkRequest) (response.WorkResponse, error)
	GetById(id string) (response.WorkResponse, error)
	GetAllWork(page, size string) ([]response.WorkResponse, string, error)
}
