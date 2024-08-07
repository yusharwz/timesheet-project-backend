package repository

import (
	"final-project-enigma/entity"
	"gorm.io/gorm"
)

type WorkRepository interface {
	CreateWork(work entity.Work) (entity.Work, error)
	UpdateWork(work entity.Work) (entity.Work, error)
	DeleteWork(id string) error
	GetById(id string) (entity.Work, error)
	GetAllWork(spec []func(db *gorm.DB) *gorm.DB) ([]entity.Work, string, error)
}
