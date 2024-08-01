package controller

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/service/impl"
	"final-project-enigma/utils"

	"github.com/gin-gonic/gin"
)

type WorkController struct{}

var workService = impl.NewWorkService()

func NewWorkController(g *gin.RouterGroup) {
	controller := new(WorkController)
	workGroup := g.Group("/admin/works")
	{
		workGroup.POST("/", controller.CreateWork)
		workGroup.PUT("/:id", controller.UpdateWork)
		workGroup.DELETE("/:id", controller.DeleteWork)
		workGroup.GET("/", controller.GetAllWork)
		workGroup.POST("/", controller.CreateWork)
		workGroup.GET("/:id", controller.GetById)
	}
}

func (WorkController) CreateWork(c *gin.Context) {
	var workRequest request.WorkRequest
	err := c.ShouldBindJSON(&workRequest)
	if err != nil {
		validationError := utils.GetValidationError(err)
		response.NewResponseBadRequest(c, validationError, "Could not parse request", "", "")
		return
	}

	result, err := workService.CreateWork(workRequest)
	if err != nil {
		response.NewResponseError(c, err.Error(), "", "")
		return
	}

	response.NewResponseCreated(c, result, "Created new work successfully", "", "")

}

func (WorkController) GetById(c *gin.Context) {
	id := c.Param("id")
	result, err := workService.GetById(id)
	if err != nil {
		response.NewResponseError(c, err.Error(), "", "")
		return
	}
	response.NewResponseSuccess(c, result, "Success fetch work data", "", "")
}

func (*WorkController) UpdateWork(c *gin.Context) {
	id := c.Param("id")

	var request request.WorkRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		validationError := utils.GetValidationError(err)
		response.NewResponseBadRequest(c, validationError, "Invalid request", "WORK", "400")
		return
	}

	result, err := workService.UpdateWork(id, request)
	if err != nil {
		response.NewResponseError(c, err.Error(), "", "")
	}

	response.NewResponseSuccess(c, result, "Work updated successfully", "WORK", "200")
}

func (WorkController) GetAllWork(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	results, total, err := workService.GetAllWork(page, size)
	if err != nil {
		response.NewResponseError(c, err.Error(), "", "")
		return
	}
	response.NewResponseSuccessPaging(c, results, "Success fetch all work data", "", "", page, size, total)
}

func (WorkController) DeleteWork(c *gin.Context) {
	id := c.Param("id")

	err := workService.DeleteWork(id)
	if err != nil {
		response.NewResponseError(c, err.Error(), "WORK", "400")
	}

	response.NewResponseSuccess(c, nil, "Work deleted successfully", "WORK", "200")
}
