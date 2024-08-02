package config

import (
	"database/sql"
	"final-project-enigma/dto"
	"final-project-enigma/entity"
	"final-project-enigma/helper"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
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
	initRoles()

	logger.Info().Msg("Initializing admin account")
	initAdmin(in.AdminConfig.Email, in.AdminConfig.Password)

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
func initRoles() {
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
	var existsRole entity.Role
	for _, role := range roles {
		DB.Where("role_name = ?", role.RoleName).First(&existsRole)
		if existsRole.ID == "" {
			DB.Create(&role)
		}
	}
}

// init admin
func initAdmin(email, password string) {
	var adminAccount entity.Account
	DB.Where("email = ? AND password = ?", email, password).First(&adminAccount)
	if adminAccount.ID != "" {
		adminAccount.Email = email
		adminAccount.Password = password
		DB.Save(adminAccount)
		return
	}
	var adminRole entity.Role
	DB.Where("role_name = ?", "admin").First(&adminRole)

	adminUserID := uuid.NewString()

	var err error
	password, err = helper.HashPassword(password)
	if err != nil {
		log.Info().Msg("failed to hash password")
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
	DB.Create(&adminUser)
	return
}
