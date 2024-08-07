package entity

type StatusTimeSheet struct {
	ID         string `gorm:"primaryKey"`
	StatusName string `gorm:"not null"`
	TimeSheets []TimeSheet
}
