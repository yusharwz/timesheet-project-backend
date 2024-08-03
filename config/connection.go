package config

import (
	"database/sql"
	"errors"
	"final-project-enigma/dto"
	"final-project-enigma/entity"
	"final-project-enigma/helper"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(in dto.ConfigData, logger zerolog.Logger) (*gorm.DB, error) {

	logger.Info().Msg("Trying Connect to DB")

	err := autoCreateDb(in, logger)
	if err != nil {
		return nil, err
	}

	var dsn = fmt.Sprintf("host=%s user= %s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", in.DbConfig.Host, in.DbConfig.User, in.DbConfig.Pass, in.DbConfig.Database, in.DbConfig.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to open database")
		return nil, err
	}

	DB = db
	err = db.AutoMigrate(
		&entity.Role{},
		&entity.StatusTimeSheet{},
		&entity.User{},
		&entity.Account{},
		&entity.Work{},
		&entity.TimeSheet{},
		&entity.TimeSheetDetail{},
	)

	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to migrate database")
	}

	logger.Info().Msg("Successfully Connected to DB")

	logger.Info().Msg("Initializing table role")
	initRoles(logger)

	logger.Info().Msg("Initializing admin account")
	err = initAdmin(in.AdminConfig.Email, in.AdminConfig.Password)
	if err != nil {
		logger.Info().Msg(err.Error())
	} else {
		logger.Info().Msg("admin account successfully initialized")
	}

	return db, nil
}

// auto create DB
func autoCreateDb(config dto.ConfigData, logger zerolog.Logger) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DbConfig.Host, config.DbConfig.User, config.DbConfig.Pass, config.DbConfig.DbPort)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database server")
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", config.DbConfig.Database))
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "42P04" { // 42P04 is the error code for "database already exists"
				logger.Info().Msg("Database already exists, proceeding without creating it")
				return nil
			}
		}
		logger.Fatal().Err(err).Msg("Failed to create database")
		return err
	} else {
		logger.Info().Msg("Database created successfully")
	}

	return nil
}

// init roles
func initRoles(logger zerolog.Logger) {
	roles := []entity.Role{
		{
			ID:       uuid.NewString(),
			RoleName: "admin",
		},
		{
			ID:       uuid.NewString(),
			RoleName: "user",
		},
		{
			ID:       uuid.NewString(),
			RoleName: "manager",
		},
		{
			ID:       uuid.NewString(),
			RoleName: "benefit",
		},
	}
	for _, role := range roles {
		var existsRole entity.Role
		result := DB.Where("role_name = ?", role.RoleName).First(&existsRole)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				DB.Create(&role)
				logger.Info().Msg(fmt.Sprintf("Role %s created", role.RoleName))
			} else {
				logger.Info().Msg(fmt.Sprintf("failed to execute query %s", result.Error))
			}
		} else {
			logger.Info().Msg(fmt.Sprintf("Role %s already exists, proceeding without creating it", role.RoleName))
		}
	}
}

// init admin
func initAdmin(email, password string) error {
	var adminAccount entity.Account
	err := DB.Where("email = ?", email).First(&adminAccount).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return createAdminIfNotFound(email, password)
	}
	password, err = helper.HashPassword(password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	adminAccount.Email = email
	adminAccount.Password = password
	DB.Save(adminAccount)
	return errors.New("admin account already initialized")
}

func createAdminIfNotFound(email, password string) error {
	var adminRole entity.Role
	DB.Where("role_name = ?", "admin").First(&adminRole)

	adminUserID := uuid.NewString()

	var err error
	password, err = helper.HashPassword(password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	newAdminAccount := entity.Account{
		Base:     entity.Base{ID: adminUserID},
		Email:    email,
		Password: password,
		IsActive: true,
		RoleID:   adminRole.ID,
		UserID:   adminUserID,
	}
	adminUser := entity.User{
		Base:    entity.Base{ID: uuid.NewString()},
		Name:    "Admin",
		Account: newAdminAccount,
	}
	result := DB.Create(&adminUser)
	if result.Error != nil {
		return errors.New("admin account failed to create")
	}
	return nil
}
