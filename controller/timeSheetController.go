package controller

import (
	"net/http"
	"strconv"
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/helper"
	"timesheet-app/middleware"
	"timesheet-app/service"
	"timesheet-app/service/impl"
	"timesheet-app/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TimeSheetController struct{}

var timeSheetService service.TimeSheetService = impl.NewTimeSheetService()

func NewTimesheetController(g *gin.RouterGroup) {
	controller := &TimeSheetController{}
	timesheetGroup := g.Group("/timesheets", middleware.JwtAuthWithRoles("user"))
	{
		timesheetGroup.GET("/:id", controller.GetTimeSheetByID)
		timesheetGroup.GET("", controller.GetAllTimeSheets)
		timesheetGroup.POST("/", controller.CreateTimeSheet)
		timesheetGroup.PUT("/:id", controller.UpdateTimeSheet)
		timesheetGroup.DELETE("/:id", controller.DeleteTimeSheet)
		timesheetGroup.PUT("/:id/submit", controller.SubmitTimeSheet)
	}
	managerGroup := g.Group("/manager", middleware.JwtAuthWithRoles("manager"))
	{
		managerGroup.POST("/approve/timesheets/:id", controller.ApproveManagerTimeSheet)
		managerGroup.POST("/reject/timesheets/:id", controller.RejectManagerTimeSheet)
	}
	benefitGroup := g.Group("/benefit", middleware.JwtAuthWithRoles("benefit"))
	{
		benefitGroup.POST("/approve/timesheets/:id", controller.ApproveBenefitTimeSheet)
		benefitGroup.POST("/reject/timesheets/:id", controller.RejectBenefitTimeSheet)
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
	year := c.Query("year")
	period := c.Query("period")
	userId := c.Query("userId")
	status := c.Query("status")
	name := c.Query("name")

	var err error
	var totalRows string
	var totalPage string
	var parsedPeriod []string
	var results *[]response.TimeSheetResponse

	if period != "" {
		parsedPeriod, err = helper.ParsePeriod(period)
		if err != nil {
			validation := utils.GetValidationError(err)
			response.NewResponseBadRequest(c, validation)
			return
		}
	}

	if year != "" {
		_, err = strconv.Atoi(year)
		if err != nil {
			validation := utils.GetValidationError(err)
			response.NewResponseBadRequest(c, validation)
			return
		}
	}

	results, totalRows, totalPage, err = timeSheetService.GetAllTimeSheets(paging, rowsPerPage, year, userId, status, name, parsedPeriod)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccessPaging(c, results, paging, rowsPerPage, totalRows, totalPage)
}

func (TimeSheetController) ApproveManagerTimeSheet(c *gin.Context) {
	id := c.Param("id")
	authHeader := c.GetHeader("Authorization")
	userId, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	err = timeSheetService.ApproveManagerTimeSheet(id, userId)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccess(c, "Approved by manager")
}

func (TimeSheetController) RejectManagerTimeSheet(c *gin.Context) {
	id := c.Param("id")
	authHeader := c.GetHeader("Authorization")
	userId, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	err = timeSheetService.RejectManagerTimeSheet(id, userId)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccess(c, "Rejected by manager")
}

func (TimeSheetController) ApproveBenefitTimeSheet(c *gin.Context) {
	id := c.Param("id")
	authHeader := c.GetHeader("Authorization")
	userId, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	err = timeSheetService.ApproveBenefitTimeSheet(id, userId)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccess(c, "Approved by benefit")
}

func (TimeSheetController) RejectBenefitTimeSheet(c *gin.Context) {
	id := c.Param("id")
	authHeader := c.GetHeader("Authorization")
	userId, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	err = timeSheetService.RejectBenefitTimeSheet(id, userId)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccess(c, "Rejected by benefit")
}

func (TimeSheetController) SubmitTimeSheet(c *gin.Context) {
	timeSheetID := c.Param("id")

	err := timeSheetService.UpdateTimeSheetStatus(timeSheetID)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccess(c, "timesheet submitted")
}
