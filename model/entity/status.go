package entity

type Status struct {
	ID       string `gorm:"type:varchar(255);primaryKey" json:"id"`
	IsActive bool   `gorm:"type:boolean" json:"is_active"`
	Account  Account
}
