package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/env"
	"github.com/iDukeLu/believer/core/util"
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

func Cors(cors *env.Cors) gin.HandlerFunc {
	return func(c *gin.Context) {
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			c.Header("Access-Control-Allow-Origin", util.GetStringDefault(cors.AllowOrigin, defaultAllowOrigin))
			c.Header("Access-Control-Allow-Methods", util.GetStringDefault(cors.AllowMethods, defaultAllowMethods))
			c.Header("Access-Control-Allow-Headers", util.GetStringDefault(cors.AllowHeaders, defaultAllowHeaders))
			c.Header("Access-Control-Allow-Credentials", util.GetStringDefault(cors.AllowCredentials, defaultAllowCredentials))
			c.Header("Access-Control-Expose-Headers", util.GetStringDefault(cors.ExposeHeaders, defaultExposeHeaders))
			c.Header("Access-Control-Max-Age", util.GetStringDefault(cors.MaxAge, defaultMaxAge))
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
