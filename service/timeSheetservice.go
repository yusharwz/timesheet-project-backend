package service

import (
    "errors" 
    "final-project-enigma/entity"
    "final-project-enigma/repository"
	"fmt"
)

type TimeSheetService interface {
    CreateTimeSheet(ts *entity.TimeSheet) error
    UpdateTimeSheet(ts *entity.TimeSheet) error
    DeleteTimeSheet(id string) error
    GetTimeSheetByID(id string) (*entity.TimeSheet, error)
    GetAllTimeSheets() (*[]entity.TimeSheet, error)
    ApproveManagerTimeSheet(id string, userID string) error
    RejectManagerTimeSheet(id string, userID string) error
    ApproveBenefitTimeSheet(id string, userID string) error
    RejectBenefitTimeSheet(id string, userID string) error
}

type timeSheetService struct {
    repo repository.TimeSheetRepository
}

func NewTimeSheetService(repo repository.TimeSheetRepository) TimeSheetService {
    return &timeSheetService{repo: repo}
}

func (s *timeSheetService) CreateTimeSheet(ts *entity.TimeSheet) error {
	fmt.Println("Error disini")
    return s.repo.CreateTimeSheet(ts)
}

func (s *timeSheetService) UpdateTimeSheet(ts *entity.TimeSheet) error {
    existingTs, err := s.repo.GetTimeSheetByID(ts.ID)
    if err != nil {
        return err
    }

    if existingTs.ConfirmedManagerBy != "" || existingTs.ConfirmedBenefitBy != "" {
        return errors.New("timesheet cannot be updated as it has been approved or rejected")
    }

    return s.repo.UpdateTimeSheet(ts)
}

func (s *timeSheetService) DeleteTimeSheet(id string) error {
    existingTs, err := s.repo.GetTimeSheetByID(id)
    if err != nil {
        return err
    }

    if existingTs.ConfirmedManagerBy != "" || existingTs.ConfirmedBenefitBy != "" {
        return errors.New("timesheet cannot be deleted as it has been approved or rejected")
    }

    return s.repo.DeleteTimeSheet(id)
}

func (s *timeSheetService) GetTimeSheetByID(id string) (*entity.TimeSheet, error) {
    return s.repo.GetTimeSheetByID(id)
}

func (s *timeSheetService) GetAllTimeSheets() (*[]entity.TimeSheet, error) {
    return s.repo.GetAllTimeSheets()
}

func (s *timeSheetService) ApproveManagerTimeSheet(id string, userID string) error {
    return s.repo.ApproveManagerTimeSheet(id, userID)
}

func (s *timeSheetService) RejectManagerTimeSheet(id string, userID string) error {
    return s.repo.RejectManagerTimeSheet(id, userID)
}

func (s *timeSheetService) ApproveBenefitTimeSheet(id string, userID string) error {
    return s.repo.ApproveBenefitTimeSheet(id, userID)
}

func (s *timeSheetService) RejectBenefitTimeSheet(id string, userID string) error {
    return s.repo.RejectBenefitTimeSheet(id, userID)
}
