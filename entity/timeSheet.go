package entity

import "gorm.io/gorm"

type TimeSheet struct {
	Base
	ID                 string            `json:"id"`
	ConfirmedManagerBy string            `json:"id_manager"`
	ConfirmedBenefitBy string            `json:"id_benefit"`
	StatusTimeSheetID  string            `json:"status_time_sheet"`
	UserID             string            `json:"user_id"`
	TimeSheetDetails   []TimeSheetDetail `json:"time_sheet_details"`
	DeletedAt          gorm.DeletedAt    `gorm:"index"`
}

// menambahkan fitur deletedat pada timesheet agar timesheet bisa soft deleted
