package helper

import (
	"gorm.io/gorm"
)

func SelectByPeriod(year, start, end string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		startDate := year + "-" + start + "-21 00:00:00"
		endDate := year + "-" + end + "-19 23:59:59"
		return db.Where("updated_at BETWEEN ? AND ?", startDate, endDate)
	}
}

func SelectByUserId(userId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}
}

func SelectByStatus(idStatus string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status_time_sheet_id = ?", idStatus)
	}
}
