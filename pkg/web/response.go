package web

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"access_token"`
}

func NewSuccessResponse(
	context *gin.Context,
	status int,
	data interface{},
) {
	context.JSON(status,
		SuccessResponse{
			status,
			data,
		})
}

func NewErrorResponse(
	context *gin.Context,
	status int,
	message string,
) {
	context.JSON(status,
		ErrorResponse{
			status,
			message,
		})
}

func NewLoginResponse(
	context *gin.Context,
	status int,
	data LoginResponse,
) {
	context.JSON(status, data)
}
