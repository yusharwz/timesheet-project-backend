package controller

import (
	"final-project-enigma/entity"
	"final-project-enigma/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TimeSheetController struct {
	service service.TimeSheetService
}

func NewTimesheetController(g *gin.RouterGroup) {
	controller := new(TimeSheetController)
	timesheetGroup := g.Group("/timesheet")
	{
		timesheetGroup.POST("/", controller.CreateTimeSheet)
		timesheetGroup.PUT("/:id", controller.UpdateTimeSheet)
		timesheetGroup.DELETE("/:id", controller.DeleteTimeSheet)
		timesheetGroup.GET("/:id", controller.GetTimeSheetByID)
		timesheetGroup.GET("/", controller.GetAllTimeSheets)

	}
}

func (ctrl *TimeSheetController) CreateTimeSheet(c *gin.Context) {
	var request entity.TimeSheet
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.CreateTimeSheet(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, request)
}

func (ctrl *TimeSheetController) UpdateTimeSheet(c *gin.Context) {
	id := c.Param("id")

	var request entity.TimeSheet
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.ID = id
	err := ctrl.service.UpdateTimeSheet(&request)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

func (c *TimeSheetController) DeleteTimeSheet(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Printf("Received request to delete time sheet with ID: %s", id)

	err := c.service.DeleteTimeSheet(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (ctrl *TimeSheetController) GetTimeSheetByID(c *gin.Context) {
	id := c.Param("id")

	timeSheet, err := ctrl.service.GetTimeSheetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, timeSheet)
}

func (ctrl *TimeSheetController) GetAllTimeSheets(c *gin.Context) {
	timeSheets, err := ctrl.service.GetAllTimeSheets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, timeSheets)
}

func (ctrl *TimeSheetController) ApproveManagerTimeSheet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Query("user_id")

	err := ctrl.service.ApproveManagerTimeSheet(id, userID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Manager approved"})
}

func (ctrl *TimeSheetController) RejectManagerTimeSheet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Query("user_id")

	err := ctrl.service.RejectManagerTimeSheet(id, userID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Manager rejected"})
}

func (ctrl *TimeSheetController) ApproveBenefitTimeSheet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Query("user_id")

	err := ctrl.service.ApproveBenefitTimeSheet(id, userID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Benefit approved"})
}

func (ctrl *TimeSheetController) RejectBenefitTimeSheet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Query("user_id")

	err := ctrl.service.RejectBenefitTimeSheet(id, userID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Benefit rejected"})
}
