package entity

type Work struct {
	Base
	Description      string `gorm:"not null" json:"description"`
	Fee              int    `gorm:"not null" json:"fee"`
	TimeSheetDetails []TimeSheetDetail
}
