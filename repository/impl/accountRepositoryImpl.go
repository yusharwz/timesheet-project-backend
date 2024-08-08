package impl

import (
	"context"
	"errors"
	"os"
	"timesheet-app/config"
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/entity"
	"timesheet-app/helper"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/rs/zerolog/log"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (AccountRepository) AccountActivation(email, password string) error {

	result := config.DB.Model(&entity.Account{}).Where("email = ? AND password = ?", email, password).Update("is_active", true)
	if result.Error != nil {
		return errors.New("failed to activate account")
	}

	return nil
}

func (AccountRepository) EditAccount(req request.EditAccountRequest) error {

	var account entity.Account
	var user entity.User

	if err := config.DB.Where("user_id = ?", req.UserID).First(&account).Error; err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	if err := config.DB.Where("id = ?", req.UserID).First(&user).Error; err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	if req.Email != "" && req.Email != account.Email {
		var existingAccount entity.Account
		if err := config.DB.Where("email = ?", req.Email).First(&existingAccount).Error; err == nil {
			log.Error()
			return errors.New("email already in use")
		}
		account.Email = req.Email
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		user.PhoneNumber = req.Phone
	}

	if err := config.DB.Save(&account).Error; err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	if err := config.DB.Save(&user).Error; err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

func (AccountRepository) UserUploadSignatureIMG(req request.UploadImagesRequest) (response.UploadImageResponse, error) {
	cldService, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	ctx := context.Background()

	var resp response.UploadImageResponse

	uploadResponse, err := cldService.Upload.Upload(ctx, req.SignatureImage, uploader.UploadParams{})
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	resp.ImageURL = uploadResponse.SecureURL

	err = config.DB.Model(&entity.User{}).Where("id = ?", req.UserID).Update("signature", resp.ImageURL).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	return resp, nil
}

func (repo AccountRepository) ChangePassword(id string, req request.ChangePasswordRequest) error {

	var account entity.Account
	if err := config.DB.Where("user_id = ? AND is_active = ?", id, true).First(&account).Error; err != nil {
		log.Error().Msg(err.Error())
		return errors.New("failed to change password")
	}

	hashedPassword, err := helper.HashPassword(req.NewPassword)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	account.Password = hashedPassword

	if err := config.DB.Save(&account).Error; err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

func (AccountRepository) GetAccountDetailByUserID(userID string) (entity.Account, entity.User, error) {
	var account entity.Account
	var user entity.User

	if err := config.DB.Where("user_id = ?", userID).First(&account).Error; err != nil {
		log.Error().Msg(err.Error())
		return account, user, err
	}

	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		log.Error().Msg(err.Error())
		return account, user, err
	}

	return account, user, nil
}
