package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	cErr "bookmark/internal/pkg/error"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
	c.Abort()
}

func Fail(c *gin.Context, httpCode int, code int, message string) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
	c.Abort()
}

func ServerError(c *gin.Context, err interface{}) {
	msg := "Internal Server Error"
	if gin.Mode() != gin.ReleaseMode {
		if _, ok := err.(error); ok {
			msg = err.(error).Error()
		}
	}
	FailByErr(c, cErr.InternalServerError(msg))
}

func FailByErr(c *gin.Context, err error) {
	v, ok := err.(*cErr.Error)

	if ok {
		Fail(c, v.HttpCode(), v.ErrorCode(), v.Error())
	} else {
		Fail(c, http.StatusBadRequest, cErr.DEFAULT_ERROR, err.Error())
	}
}
