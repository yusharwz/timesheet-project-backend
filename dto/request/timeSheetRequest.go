package request

import "time"

type TimeSheetRequest struct {
	UserID             string                   `json:"user_id"`
	StatusTimeSheetID  string                   `json:"status_time_sheet_id"`
	ConfirmedManagerBy string                   `json:"id_manager"`
	ConfirmedBenefitBy string                   `json:"id_benefit"`
	TimeSheetDetails   []TimeSheetDetailRequest `json:"time_sheet_details"`
}

type TimeSheetDetailRequest struct {
	Date      time.Time `json:"date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	WorkID    string    `json:"work_id"`
}
