package handler

import (
	"alpha/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseErr struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data"`
	Errors  map[string]string `json:"errors"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendBadResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusBadRequest
	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendBadResponseErrors(c *gin.Context, err error, data interface{}, Errors map[string]string) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusBadRequest
	c.JSON(http.StatusBadRequest, ResponseErr{
		Code:    code,
		Message: message,
		Data:    data,
		Errors:  Errors,
	})
}
