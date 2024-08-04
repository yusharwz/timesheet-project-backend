package repository

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
)

type AccountRepository interface {
	AccountActivation(email, password string) error
	EditAccount(req request.EditAccountRequest) error
	UserUploadSignatureIMG(req request.UploadImagesRequest) (response.UploadImageResponse, error)
	ChangePassword(id string, req request.ChangePasswordRequest) error
	GetAccountDetailByUserID(userID string) (entity.Account, entity.User, error)
}
