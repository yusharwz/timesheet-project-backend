package controller

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
	"final-project-enigma/middleware"
	"final-project-enigma/service"
	"final-project-enigma/service/impl"
	"final-project-enigma/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TimeSheetController struct{}

var timeSheetService service.TimeSheetService = impl.NewTimeSheetService()

func NewTimesheetController(g *gin.RouterGroup) {
	controller := new(TimeSheetController)
	timesheetGroup := g.Group("/timesheets", middleware.JwtAuthWithRoles("user"))
	{
		timesheetGroup.POST("/", middleware.JwtAuthWithRoles("user"), controller.CreateTimeSheet)
		timesheetGroup.PUT("/:id", middleware.JwtAuthWithRoles("user"), controller.UpdateTimeSheet)
		timesheetGroup.DELETE("/:id", middleware.JwtAuthWithRoles("user"), controller.DeleteTimeSheet)

	}
	g.GET("/timesheets/:id", controller.GetTimeSheetByID)
	g.GET("/timesheets", controller.GetAllTimeSheets)
	managerGroup := g.Group("/manager", middleware.JwtAuthWithRoles("manager"))
	{
		managerGroup.POST("approve/timesheets/:id", controller.ApproveManagerTimeSheet)
		managerGroup.POST("reject/timesheets/:id", controller.RejectManagerTimeSheet)
	}
	benefitGroup := g.Group("benefit", middleware.JwtAuthWithRoles("benefit"))
	{
		benefitGroup.POST("approve/timesheets/:id", controller.ApproveBenefitTimeSheet)
		benefitGroup.POST("reject/timesheets/:id", controller.RejectBenefitTimeSheet)
	}
}

func (TimeSheetController) CreateTimeSheet(c *gin.Context) {
	var req request.TimeSheetRequest
	authHeader := c.GetHeader("Authorization")

	if err := c.ShouldBindJSON(&req); err != nil {
		validation := utils.GetValidationError(err)
		response.NewResponseBadRequest(c, validation)
		return
	}

	res, err := timeSheetService.CreateTimeSheet(req, authHeader)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseCreated(c, res)
}

func (TimeSheetController) UpdateTimeSheet(c *gin.Context) {
	id := c.Param("id")

	var req entity.TimeSheet
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = id
	err := timeSheetService.UpdateTimeSheet(&req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}

func (TimeSheetController) DeleteTimeSheet(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Printf("Received request to delete time sheet with ID: %s", id)

	err := timeSheetService.DeleteTimeSheet(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (TimeSheetController) GetTimeSheetByID(c *gin.Context) {
	id := c.Param("id")

	timeSheet, err := timeSheetService.GetTimeSheetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, timeSheet)
}

func (TimeSheetController) GetAllTimeSheets(c *gin.Context) {
	timeSheets, err := timeSheetService.GetAllTimeSheets()
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccess(c, timeSheets)
}

func (TimeSheetController) ApproveManagerTimeSheet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Query("user_id")

	err := timeSheetService.ApproveManagerTimeSheet(id, userID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Manager approved"})
}

func (TimeSheetController) RejectManagerTimeSheet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Query("user_id")

	err := timeSheetService.RejectManagerTimeSheet(id, userID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Manager rejected"})
}

func (TimeSheetController) ApproveBenefitTimeSheet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Query("user_id")

	err := timeSheetService.ApproveBenefitTimeSheet(id, userID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Benefit approved"})
}

func (TimeSheetController) RejectBenefitTimeSheet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Query("user_id")

	err := timeSheetService.RejectBenefitTimeSheet(id, userID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Benefit rejected"})
}
