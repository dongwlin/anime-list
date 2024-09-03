package api

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type Pagination struct {
	TotalItems  int64 `json:"total_items,omitempty"`
	TotalPages  int64 `json:"total_pages,omitempty"`
	PerPage     int64 `json:"per_page"`
	CurrentPage int64 `json:"current_page"`
}

func SuccessResponse(ctx *gin.Context, httpStatus int, message string, data interface{}, meta interface{}) {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	ctx.JSON(httpStatus, response)
}

func ErrorResponse(ctx *gin.Context, httpStatus int, message string, errors interface{}) {
	response := Response{
		Success: false,
		Message: message,
		Errors:  errors,
	}
	ctx.JSON(httpStatus, response)
}
