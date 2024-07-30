package entity

type StatusTimeSheet struct {
	ID         string `gorm:"primaryKey" json:"id"`
	StatusName string `gorm:"not null" json:"status_name"`
	TimeSheets []TimeSheet
}
