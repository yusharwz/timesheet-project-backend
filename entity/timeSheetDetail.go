package entity

import "time"

type TimeSheetDetail struct {
	Base
	Date        time.Time `json:"date"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	TimeSheetID string    `json:"time_sheet_id"`
	WorkID      string    `json:"work_id"`
}
