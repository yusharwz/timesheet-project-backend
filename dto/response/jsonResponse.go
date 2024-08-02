package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type (
	jsonResponse struct {
		Code    int         `json:"responseCode"`
		Message string      `json:"responseMessage,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}

	jsonResponseWithPaging struct {
		Code    int         `json:"responseCode"`
		Message string      `json:"responseMessage"`
		Data    interface{} `json:"data,omitempty"`
		Paging  interface{} `json:"paging,omitempty"`
	}

	jsonErrorResponse struct {
		Code    int    `json:"responseCode"`
		Message string `json:"responseMessage"`
		Error   string `json:"error,omitempty"`
	}

	ValidationField struct {
		FieldName int    `json:"field"`
		Message   string `json:"message"`
	}

	jsonBadRequestResponse struct {
		Code             int               `json:"responseCode"`
		Message          string            `json:"responseMessage"`
		ErrorDescription []ValidationField `json:"error_description,omitempty"`
	}

	PagingInfo struct {
		Page      string `json:"page,omitempty"`
		Size      string `json:"size,omitempty"`
		TotalData string `json:"totalData,omitempty"`
	}
)

func NewResponseSuccessPaging(c *gin.Context, result interface{}, message, page, size, totalData string) {
	c.JSON(http.StatusOK, jsonResponseWithPaging{
		Code:    http.StatusOK,
		Message: message,
		Data:    result,
		Paging: PagingInfo{
			Page:      page,
			Size:      size,
			TotalData: totalData,
		},
	})
}

func NewResponseSuccess(c *gin.Context, result interface{}, message string) {
	c.JSON(http.StatusOK, jsonResponse{
		Code:    http.StatusOK,
		Message: message,
		Data:    result,
	})
}

func NewResponseCreated(c *gin.Context, result interface{}, message string) {
	c.JSON(http.StatusCreated, jsonResponse{
		Code:    http.StatusCreated,
		Message: message,
		Data:    result,
	})
}

func NewResponseBadRequest(c *gin.Context, validationField []ValidationField, message string) {
	c.JSON(http.StatusBadRequest, jsonBadRequestResponse{
		Code:             http.StatusBadRequest,
		Message:          message,
		ErrorDescription: validationField,
	})
}

func NewResponseError(c *gin.Context, err string) {
	log.Error().Msg(err)
	c.JSON(http.StatusInternalServerError, jsonErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Error:   err,
	})
}

func NewResponseForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, jsonErrorResponse{
		Code:    http.StatusForbidden,
		Message: message,
	})
}

func NewResponseUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, jsonErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: message,
	})
}
