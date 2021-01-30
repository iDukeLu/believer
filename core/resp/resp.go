package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{http.StatusOK, msg, data})
}

func Failure(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{http.StatusInternalServerError, msg, data})
}

func DefaultSuccess(c *gin.Context, data interface{}) {
	Success(c, success, data)
}

func DefaultFailure(c *gin.Context, data interface{}) {
	Failure(c, failure, data)
}
