package entity

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	Base
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	IsActive     bool   `gorm:"not null"`
	LoginChances int
	LoginTime    time.Time
	RoleID       string
	UserID       string
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	if a.LoginChances == 0 {
		a.LoginChances = 3
	}
	if a.LoginTime.IsZero() {
		a.LoginTime = time.Now()
	}
	return
}
