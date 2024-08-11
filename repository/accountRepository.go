package repository

import (
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/entity"
)

type AccountRepository interface {
	AccountActivation(email, password string) error
	EditAccount(req request.EditAccountRequest) error
	UserUploadSignatureIMG(req request.UploadImagesRequest) (response.UploadImageResponse, error)
	ChangePassword(id string, req request.ChangePasswordRequest) error
	GetAccountDetailByUserID(userID string) (entity.Account, entity.User, error)
	ForgetPassword(req request.ForgetPasswordRequest) error
}
