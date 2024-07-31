package response

import "time"

type TimeSheetResponse struct {
	ID                 string                 `json:"id"`
	UserID             string                 `json:"user_id"`
	StatusTimeSheetID  string                 `json:"status_time_sheet_id"`
	ConfirmedManagerBy string                `json:"confirmed_manager_by,omitempty"`
	ConfirmedBenefitBy string                `json:"confirmed_benefit_by,omitempty"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
	TimeSheetDetails   []TimeSheetDetailResponse `json:"time_sheet_details"`
}

type TimeSheetDetailResponse struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	WorkID      string    `json:"work_id"`
	TimeSheetID string    `json:"time_sheet_id"`
}