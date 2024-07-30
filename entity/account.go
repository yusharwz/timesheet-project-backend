package entity

type Account struct {
	Base
	Email    string `gorm:"unique;not null" json:"email"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	IsActive bool   `gorm:"not null" json:"is_active"`
	RoleID   string `json:"role_id"`
	UserID   string `json:"user_id"`
}
