package request

import "time"

type TimeSheetRequest struct {
	TimeSheetDetails []TimeSheetDetailRequest `json:"timeSheetDetails" binding:"required"`
}

type UpdateTimeSheetRequest struct {
	ID               string                         `json:"id"`
	TimeSheetDetails []UpdateTimeSheetDetailRequest `json:"timeSheetDetails" binding:"required"`
}

type TimeSheetDetailRequest struct {
	Date      time.Time `json:"date" binding:"required"`
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
	WorkID    string    `json:"workId" binding:"required"`
}

type UpdateTimeSheetDetailRequest struct {
	ID        string    `json:"id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
	WorkID    string    `json:"workId" binding:"required"`
}

type UpdateTimeSheetStatusRequest struct {
	TimeSheetID string
}
