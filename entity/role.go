package entity

type Role struct {
	ID       string `gorm:"primaryKey" json:"id"`
	RoleName string `gorm:"not null" json:"role_name"`
	Accounts []Account
}
