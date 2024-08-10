package impl

import (
	"errors"
	"time"
	"timesheet-app/config"
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/entity"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (AuthRepository) Register(user entity.User, account entity.Account) (entity.User, entity.Account, error) {

	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Error()
			tx.Rollback()
		}
	}()

	var existingAccount entity.Account
	if err := tx.Where("email = ?", account.Email).First(&existingAccount).Error; err == nil {
		log.Error()
		tx.Rollback()
		return user, account, errors.New("email already in use")
	}

	user.Account = account
	if err := tx.Create(&user).Error; err != nil {
		log.Error().Msg(err.Error())
		tx.Rollback()
		return user, account, errors.New("failed to create account")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error().Msg(err.Error())
		tx.Rollback()
		return user, account, errors.New("transaction commit failed")
	}

	return user, account, nil
}

func (AuthRepository) Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error) {
	var account entity.Account
	var user entity.User
	var role entity.Role
	timeNow := time.Now()

	resultAccount := config.DB.Where("email = ?", req.Email).First(&account)
	if resultAccount.Error != nil {
		log.Error().Msg(resultAccount.Error.Error())
		if errors.Is(resultAccount.Error, gorm.ErrRecordNotFound) {
			return resp, errors.New("invalid email or password")
		}
		return resp, resultAccount.Error
	}

	if account.LoginChances == 0 {
		if account.LoginTime.After(timeNow) {
			log.Error().Msg("Account is locked. Please try again after 15 minutes")
			return resp, errors.New("account has been locked due to too many login attempts, please try again after 15 minutes on " + account.LoginTime.Format(time.RFC1123))
		} else {
			account.LoginChances = 3
			if err := config.DB.Save(&account).Error; err != nil {
				log.Error().Msg(err.Error())
				return resp, err
			}
		}
	}

	resultUser := config.DB.Where("id = ?", account.UserID).First(&user)
	if resultUser.Error != nil {
		log.Error().Msg(resultUser.Error.Error())
		if errors.Is(resultUser.Error, gorm.ErrRecordNotFound) {
			return resp, errors.New("invalid email or password")
		}
		return resp, resultUser.Error
	}

	if err := config.DB.Where("id = ?", account.RoleID).First(&role).Error; err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	if !account.IsActive {
		log.Error().Msg("Account is not active")
		return resp, errors.New("account is not active")
	}

	if !account.DeletedAt.Time.IsZero() {
		log.Error().Msg("Account has been deleted")
		return resp, errors.New("account has been deleted")
	}

	resp.HashPassword = account.Password
	resp.Email = account.Email
	resp.Name = user.Name
	resp.UserId = account.UserID
	resp.Role = role.RoleName
	resp.LoginChance = account.LoginChances
	resp.LoginTime = account.LoginTime

	return resp, nil
}

func (AuthRepository) DecrementLoginChance(email string) error {
	var account entity.Account
	timeNow := time.Now()

	result := config.DB.Where("email = ?", email).First(&account)
	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		return result.Error
	}

	if account.LoginChances == 1 {
		account.LoginTime = timeNow.Add(15 * time.Minute)
	}
	account.LoginChances--

	if err := config.DB.Save(&account).Error; err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}

func (AuthRepository) IncrementLoginChance(email string) error {
	var account entity.Account

	result := config.DB.Where("email = ?", email).First(&account)
	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		return result.Error
	}
	account.LoginChances = 3

	if err := config.DB.Save(&account).Error; err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}

func (AuthRepository) GetRoleByName(roleName string) (entity.Role, error) {
	var role entity.Role
	result := config.DB.Where("role_name = ?", roleName).First(&role)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Error()
			return entity.Role{}, errors.New("invalid role")
		}
		return entity.Role{}, result.Error
	}
	return role, nil
}

func (AuthRepository) GetRoleById(id string) (*entity.Role, error) {
	var role entity.Role
	err := config.DB.Where("id = ?", id).First(&role).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return &role, nil
}
