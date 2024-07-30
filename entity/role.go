package entity

import "github.com/google/uuid"

type Role struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	RoleName string    `gorm:"not null" json:"role_name"`
	Accounts []Account
}
