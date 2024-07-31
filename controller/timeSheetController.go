package controller

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TimeSheetController struct {
	ts service.TimeSheetService
}

func NewTimeSheetController(ts service.TimeSheetService) *TimeSheetController {
	return &TimeSheetController{ts: ts}
}

func (c *TimeSheetController) CreateTimeSheet(ctx *gin.Context) {
	var tsRequest request.TimeSheetRequest
	if err := ctx.ShouldBindJSON(&tsRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tsResponse, err := c.ts.CreateTimeSheet(&tsRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, tsResponse)
}

func (c *TimeSheetController) FindByIdTimeSheet(ctx *gin.Context) {
	id := ctx.Param("id")

	tsResponse, err := c.ts.FindByIdTimeSheet(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tsResponse)
}

func (c *TimeSheetController) FindAllTimeSheet(ctx *gin.Context) {
	tsResponses, err := c.ts.FindAllTimeSheet()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tsResponses)
}

func (c *TimeSheetController) UpdateTimeSheet(ctx *gin.Context) {
	var tsRequest request.TimeSheetRequest
	if err := ctx.ShouldBindJSON(&tsRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tsResponse, err := c.ts.UpdateTimeSheet(&tsRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tsResponse)
}

func (c *TimeSheetController) DeleteTimeSheet(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.ts.DeleteTimeSheet(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "TimeSheet deleted successfully"})
}
