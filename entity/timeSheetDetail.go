package entity

import "time"

type TimeSheetDetail struct {
	Base
	Date        time.Time
	StartTime   time.Time
	EndTime     time.Time
	TimeSheetID string
	WorkID      string
}
