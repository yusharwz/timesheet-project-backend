package controller

import (
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/middleware"
	"timesheet-app/service"
	"timesheet-app/service/impl"
	"timesheet-app/utils"

	"github.com/gin-gonic/gin"
)

type WorkController struct{}

var workService service.WorkService = impl.NewWorkService()

func NewWorkController(g *gin.RouterGroup) {
	controller := new(WorkController)
	workGroup := g.Group("/admin/works", middleware.JwtAuthWithRoles("user", "admin"))
	{
		workGroup.GET("", controller.GetAllWork)
		workGroup.POST("/", middleware.JwtAuthWithRoles("admin"), controller.CreateWork)
		workGroup.PUT("/:id", middleware.JwtAuthWithRoles("admin"), controller.UpdateWork)
		workGroup.DELETE("/:id", middleware.JwtAuthWithRoles("admin"), controller.DeleteWork)
		workGroup.GET("/:id", middleware.JwtAuthWithRoles("admin"), controller.GetById)
	}
}

func (WorkController) CreateWork(c *gin.Context) {
	var workRequest request.WorkRequest
	err := c.ShouldBindJSON(&workRequest)
	if err != nil {
		validationError := utils.GetValidationError(err)
		response.NewResponseBadRequest(c, validationError)
		return
	}

	result, err := workService.CreateWork(workRequest)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseCreated(c, result)
}

func (WorkController) GetById(c *gin.Context) {
	id := c.Param("id")
	result, err := workService.GetById(id, false)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}
	response.NewResponseSuccess(c, result)
}

func (*WorkController) UpdateWork(c *gin.Context) {
	id := c.Param("id")

	var r request.WorkRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		validationError := utils.GetValidationError(err)
		response.NewResponseBadRequest(c, validationError)
		return
	}

	result, err := workService.UpdateWork(id, r)
	if err != nil {
		response.NewResponseError(c, err.Error())
	}

	response.NewResponseSuccess(c, result)
}

func (WorkController) GetAllWork(c *gin.Context) {
	paging := c.DefaultQuery("paging", "1")
	rowsPerPage := c.DefaultQuery("rowsPerPage", "10")
	description := c.Query("description")
	results, totalRows, totalPage, err := workService.GetAllWork(paging, rowsPerPage, description)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseSuccessPaging(c, results, paging, rowsPerPage, totalRows, totalPage)
}

func (WorkController) DeleteWork(c *gin.Context) {
	id := c.Param("id")

	err := workService.DeleteWork(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
	}

	response.NewResponseSuccess(c, nil)
}
