package entity

type Account struct {
	Base
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	IsActive bool   `gorm:"not null"`
	RoleID   string
	UserID   string
}
