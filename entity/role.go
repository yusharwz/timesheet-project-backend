package entity

type Role struct {
	ID       string `gorm:"primaryKey" `
	RoleName string `gorm:"not null"`
	Accounts []Account
}
