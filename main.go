package main

import (
    "final-project-enigma/controller"
    "final-project-enigma/repository"
    "final-project-enigma/service"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"os"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "final-project-enigma/entity"
)

func main() {
    // Initialize logger
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

    // Database connection
    dsn := "user=postgres password=26527y25tw dbname=ts host=localhost port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to connect to database")
    }

    // Database migration
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
        log.Fatal().Err(err).Msg("Failed to migrate database")
    }


    repo := repository.NewTimeSheetRepository(db)
    tsService := service.NewTimeSheetService(repo)
    tsController := controller.NewTimeSheetController(tsService)


    router := gin.Default()


    router.POST("/timesheets", tsController.CreateTimeSheet)
    router.PUT("/timesheets/:id", tsController.UpdateTimeSheet)
    router.DELETE("/timesheets/:id", tsController.DeleteTimeSheet)
    router.GET("/timesheets/:id", tsController.GetTimeSheetByID)
    router.GET("/timesheets", tsController.GetAllTimeSheets)

    router.POST("/manager/approve/timesheets/:id", tsController.ApproveManagerTimeSheet)
    router.POST("/manager/reject/timesheets/:id", tsController.RejectManagerTimeSheet)

    router.POST("/benefit/approve/timesheets/:id", tsController.ApproveBenefitTimeSheet)
    router.POST("/benefit/reject/timesheets/:id", tsController.RejectBenefitTimeSheet)

  
    router.Run(":8080")
}
