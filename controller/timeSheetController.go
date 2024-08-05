package controller

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/helper"
	"final-project-enigma/middleware"
	"final-project-enigma/service"
	"final-project-enigma/service/impl"
	"final-project-enigma/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TimeSheetController struct{}

var timeSheetService service.TimeSheetService = impl.NewTimeSheetService()

func NewTimesheetController(g *gin.RouterGroup) {
	controller := &TimeSheetController{}
	timesheetGroup := g.Group("/timesheets", middleware.JwtAuthWithRoles("user"))
	{
		timesheetGroup.POST("/", controller.CreateTimeSheet)
		timesheetGroup.PUT("/:id", controller.UpdateTimeSheet)
		timesheetGroup.DELETE("/:id", controller.DeleteTimeSheet)
		timesheetGroup.PUT(":id/submit", controller.SubmitTimeSheet)
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
	authHeader := c.GetHeader("Authorization")

	var req request.UpdateTimeSheetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validation := utils.GetValidationError(err)
		response.NewResponseBadRequest(c, validation)
		return
	}

	req.ID = id
	res, err := timeSheetService.UpdateTimeSheet(req, authHeader)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccess(c, res)
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
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccess(c, timeSheet)
}

func (TimeSheetController) GetAllTimeSheets(c *gin.Context) {
	paging := c.DefaultQuery("paging", "1")
	rowsPerPage := c.DefaultQuery("rowsPerPage", "10")
	period := c.Query("period")
	userId := c.Query("userId")
	confirm := c.Query("confirm")
	status := c.Query("status")

	var err error
	var totalRows string
	var totalPage string
	var results *[]response.TimeSheetResponse

	if period != "" {
		err = helper.ParsePeriod(period)
		if err != nil {
			validation := utils.GetValidationError(err)
			response.NewResponseBadRequest(c, validation)
			return
		}
	}

	results, totalRows, totalPage, err = timeSheetService.GetAllTimeSheets(paging, rowsPerPage, period, userId, confirm, status)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccessPaging(c, results, paging, rowsPerPage, totalRows, totalPage)
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

func (TimeSheetController) SubmitTimeSheet(c *gin.Context) {

}
