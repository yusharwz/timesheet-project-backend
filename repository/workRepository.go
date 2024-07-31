package repository

import "final-project-enigma/entity"

type WorkRepository interface {
	CreateWork(work entity.Work) (entity.Work, error)
	GetById(id string) (entity.Work, error)
}
