package request

import "time"

type TimeSheetRequest struct {
	TimeSheetDetails []TimeSheetDetailRequest `json:"time_sheet_details"`
}

type TimeSheetDetailRequest struct {
	Date      time.Time `json:"date" binding:"required"`
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
	WorkID    string    `json:"workId" binding:"required"`
}
