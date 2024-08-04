package testutils

import (
	"final-project-enigma/entity" // Pastikan import path sesuai dengan proyek Anda
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate all your models
	err = db.AutoMigrate(&entity.TimeSheet{}, &entity.StatusTimeSheet{}, &entity.TimeSheetDetail{})
	if err != nil {
		return nil, nil
	}

	// Cleanup function to close the database connection
	cleanup := func() {
		sqlDB, _ := db.DB()
		err := sqlDB.Close()
		if err != nil {
			return
		}
	}

	return db, cleanup
}
