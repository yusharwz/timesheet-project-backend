package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	jsonResponse struct {
		Status Status      `json:"status"`
		Data   interface{} `json:"data,omitempty"`
	}

	jsonResponseWithPaging struct {
		Status Status      `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Paging Paging      `json:"paging,omitempty"`
	}

	jsonErrorResponse struct {
		Status Status `json:"status"`
		Data   string `json:"data,omitempty"`
	}

	ValidationField struct {
		FieldName string `json:"field"`
		Message   string `json:"message"`
	}

	jsonBadRequestResponse struct {
		Status Status            `json:"status"`
		Data   []ValidationField `json:"data,omitempty"`
	}

	Paging struct {
		Paging      string `json:"paging,omitempty"`
		RowsPerPage string `json:"rowsPerPage,omitempty"`
		TotalRows   string `json:"totalRows,omitempty"`
		TotalPage   string `json:"totalPage,omitempty"`
	}

	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func NewResponseSuccessPaging(c *gin.Context, result interface{}, paging, rowsPerPage, totalRows, totalPage string) {
	c.JSON(http.StatusOK, jsonResponseWithPaging{
		Status: Status{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data: result,
		Paging: Paging{
			Paging:      paging,
			RowsPerPage: rowsPerPage,
			TotalRows:   totalRows,
			TotalPage:   totalPage,
		},
	})
}

func NewResponseSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, jsonResponse{
		Status: Status{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data: result,
	})
}

func NewResponseCreated(c *gin.Context, result interface{}) {
	c.JSON(http.StatusCreated, jsonResponse{
		Status: Status{
			Code:    http.StatusCreated,
			Message: "Created",
		},
		Data: result,
	})
}

func NewResponseBadRequest(c *gin.Context, validationField []ValidationField) {
	c.JSON(http.StatusBadRequest, jsonBadRequestResponse{
		Status: Status{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
		},
		Data: validationField,
	})
}

func NewResponseError(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, jsonErrorResponse{
		Status: Status{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		},
		Data: err,
	})
}

func NewResponseForbidden(c *gin.Context, err string) {
	c.JSON(http.StatusForbidden, jsonErrorResponse{
		Status: Status{
			Code:    http.StatusForbidden,
			Message: "Forbidden",
		},
		Data: err,
	})
}

func NewResponseUnauthorized(c *gin.Context, err string) {
	c.JSON(http.StatusUnauthorized, jsonErrorResponse{
		Status: Status{
			Code:    http.StatusOK,
			Message: "Unauthorized",
		},
		Data: err,
	})
}
