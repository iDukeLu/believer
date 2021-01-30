package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	defaultAllowOrigin      = "*"
	defaultAllowMethods     = "*"
	defaultAllowHeaders     = "*"
	defaultAllowCredentials = "true"
	defaultExposeHeaders    = "Cache-Control, Content-Language, Content-Length, Content-Type, Date, Expires"
	defaultMaxAge           = "1800"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")

		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
