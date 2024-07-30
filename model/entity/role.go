package entity

type Role struct {
	ID       string `gorm:"type:varchar(255);primaryKey" json:"id"`
	RoleName string `json:"role_name"`
	Account  Account
}
