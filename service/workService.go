package service

import (
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
)

type WorkService interface {
	CreateWork(request request.WorkRequest) (response.WorkResponse, error)
	UpdateWork(id string, request request.WorkRequest) (response.WorkResponse, error)
	DeleteWork(id string) error
	GetById(id string, useSpec, getDeleted bool) (response.WorkResponse, error)
	GetAllWork(paging, rowsPerPage, description string) ([]response.WorkResponse, string, string, error)
}
