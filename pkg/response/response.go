package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"` // status code
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponseData struct {
	Code   int         `json:"code"`   // status code
	Err    string      `json:"error"`  // thong bao loi
	Detail interface{} `json:"detail"` //du lai return
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	if message == "" {
		message = ErrMessageDict()(code)
	}
	c.JSON(http.StatusConflict, ResponseData{
		Code:    code,
		Message: message,
	})
}
