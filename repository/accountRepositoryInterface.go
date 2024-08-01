package repository

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
)

type AccountRepository interface {
	AccountActivation(email string) error
	Login(req request.LoginAccountRequest) (response response.LoginResponse, hashedPassword string, err error)
}
