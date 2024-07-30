package config

import (
	"final-project-enigma/dto"
	"final-project-enigma/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func ConnectDb(in dto.ConfigData, logger zerolog.Logger) (*gorm.DB, error) {

	logger.Info().Msg("Trying Connect to DB")

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
	return db, nil
}
