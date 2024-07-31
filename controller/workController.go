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
