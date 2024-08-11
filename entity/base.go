package entity

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
