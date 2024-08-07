package helper

import (
	"gorm.io/gorm"
	"strings"
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

func SelectAccountByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%")
	}
}

func SelectWorkByDescription(description string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("LOWER(description) LIKE ?", "%"+strings.ToLower(description)+"%")
	}
}

func SelectUserInTimSheet(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("JOIN users ON users.id = time_sheets.user_id").Where("LOWER(users.name) LIKE ?", "%"+strings.ToLower(name)+"%")
	}
}
