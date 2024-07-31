package helper

import (
	"gorm.io/gorm"
	"strconv"
)

func Paginate(page, size string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(size)

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
