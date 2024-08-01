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
	db.AutoMigrate(&entity.TimeSheet{}, &entity.StatusTimeSheet{}, &entity.TimeSheetDetail{})

	// Cleanup function to close the database connection
	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return db, cleanup
}
