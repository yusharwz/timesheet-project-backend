package response

import (
	"time"
)

type TimeSheetResponse struct {
	ID                    string                    `json:"id"`
	CreatedAt             time.Time                 `json:"createdAt"`
	UpdatedAt             time.Time                 `json:"updatedAt"`
	StatusByManager       string                    `json:"statusByManager"`
	StatusByBenefit       string                    `json:"statusByBenefit"`
	ConfirmedManagerBy    ConfirmedByResponse       `json:"confirmedManagerBy"`
	ConfirmedBenefitBy    ConfirmedByResponse       `json:"confirmedBenefitBy"`
	UserTimeSheetResponse UserTimeSheetResponse     `json:"user"`
	TimeSheetDetails      []TimeSheetDetailResponse `json:"timeSheetDetails"`
	Total                 int                       `json:"total"`
}

type TimeSheetDetailResponse struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	WorkID      string    `json:"workId"`
	Description string    `json:"description"`
	SubTotal    int       `json:"subTotal"`
}

type ConfirmedByResponse struct {
	ID           string `json:"userId"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	SignatureUrl string `json:"signatureUrl"`
}

type UserTimeSheetResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	SignatureUrl string `json:"signatureUrl"`
}
