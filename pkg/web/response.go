package web

import (
	"fmt"
	"net/http"
	"strings"

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

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, response{Data: data})
}

// NewErrorf creates a new error with the given status code and the message
// formatted according to args and format.
func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	Response(c, status, err)
}
