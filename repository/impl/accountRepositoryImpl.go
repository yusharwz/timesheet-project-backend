package impl

import (
	"errors"
	"final-project-enigma/config"
	"final-project-enigma/dto/request"
	"final-project-enigma/entity"
	"final-project-enigma/helper"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (AccountRepository) AccountActivation(email string) error {

	result := config.DB.Model(&entity.Account{}).Where("email = ?", email).Update("is_active", true)
	if result.Error != nil {
		return errors.New("disini")
	}

	return nil
}

func (AccountRepository) EditAccount(req request.EditAccountRequest) error {

	var account entity.Account
	var user entity.User

	if err := config.DB.Where("user_id = ?", req.UserID).First(&account).Error; err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", req.UserID).First(&user).Error; err != nil {
		return err
	}

	if req.Email != "" && req.Email != account.Email {
		var existingAccount entity.Account
		if err := config.DB.Where("email = ?", req.Email).First(&existingAccount).Error; err == nil {
			return errors.New("email already in use")
		}
		account.Email = req.Email
	}

	if req.Username != "" && req.Username != account.Username {
		var existingAccount entity.Account
		if err := config.DB.Where("username = ?", req.Username).First(&existingAccount).Error; err == nil {
			return errors.New("username already in use")
		}
		account.Username = req.Username
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		user.PhoneNumber = req.Phone
	}

	if err := config.DB.Save(&account).Error; err != nil {
		return err
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo AccountRepository) ChangePassword(req request.ChangePasswordRequest) error {

	var account entity.Account
	if err := config.DB.Where("user_id = ?", req.UserID).First(&account).Error; err != nil {
		return errors.New("failed to change password")
	}

	hashedPassword, err := helper.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	account.Password = string(hashedPassword)

	if err := config.DB.Save(&account).Error; err != nil {
		return err
	}

	return nil
}

func (AccountRepository) GetAccountDetailByUserID(userID string) (entity.Account, entity.User, error) {
	var account entity.Account
	var user entity.User

	if err := config.DB.Where("user_id = ?", userID).First(&account).Error; err != nil {
		return account, user, err
	}

	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return account, user, err
	}

	return account, user, nil
}
