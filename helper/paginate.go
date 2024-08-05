package helper

import (
	"gorm.io/gorm"
	"strconv"
)

func Paginate(paging, rowsPerPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if paging <= 0 {
			paging = 1
		}

		offset := (paging - 1) * rowsPerPage
		return db.Offset(offset).Limit(rowsPerPage)
	}
}

func GetTotalPage(totalRows string, rowsPerPage int) int {
	var totalPage int
	tr, _ := strconv.Atoi(totalRows)
	rest := tr % rowsPerPage
	totalPage = tr / rowsPerPage
	if rest != 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}
