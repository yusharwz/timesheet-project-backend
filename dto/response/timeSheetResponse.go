package response

import (
	"time"
)

type TimeSheetResponse struct {
	ID                 string                    `json:"id"`
	CreatedAt          time.Time                 `json:"createdAt"`
	UpdatedAt          time.Time                 `json:"updatedAt"`
	StatusByManager    string                    `json:"statusByManager"`
	StatusByBenefit    string                    `json:"statusByBenefit"`
	ConfirmedManagerBy ConfirmedByResponse       `json:"confirmedManagerBy"`
	ConfirmedBenefitBy ConfirmedByResponse       `json:"confirmedBenefitBy"`
	TimeSheetDetails   []TimeSheetDetailResponse `json:"timeSheetDetails"`
}

type TimeSheetDetailResponse struct {
	ID        string    `json:"id"`
	Date      time.Time `json:"date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	WorkID    string    `json:"work_id"`
}

type ConfirmedByResponse struct {
	ID           string `json:"userId"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	SignatureUrl string `json:"signatureUrl"`
}
