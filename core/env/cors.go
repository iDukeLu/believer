package env

import (
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/util"
	"net/http"
)

type Cors struct {
	Enable           bool   `yaml:"enable"`
	AllowOrigin      string `yaml:"allow-origin"`
	AllowMethods     string `yaml:"allow-methods"`
	AllowHeaders     string `yaml:"allow-headers"`
	AllowCredentials string `yaml:"allow-credentials"`
	ExposeHeaders    string `yaml:"expose-headers"`
	MaxAge           string `yaml:"max-age"`
}

const (
	defaultAllowOrigin      = "*"
	defaultAllowMethods     = "*"
	defaultAllowHeaders     = "*"
	defaultAllowCredentials = "true"
	defaultExposeHeaders    = "Cache-Control, Content-Language, Content-Length, Content-Type, Date, Expires"
	defaultMaxAge           = "1800"
)

func getMergeCors(defaultCors *Cors, profileCors *Cors) Cors {
	return Cors{
		util.GetBoolDefault(defaultCors.Enable, profileCors.Enable),
		util.GetStringDefault(defaultCors.AllowOrigin, profileCors.AllowOrigin),
		util.GetStringDefault(defaultCors.AllowMethods, profileCors.AllowMethods),
		util.GetStringDefault(defaultCors.AllowHeaders, profileCors.AllowHeaders),
		util.GetStringDefault(defaultCors.AllowCredentials, profileCors.AllowCredentials),
		util.GetStringDefault(defaultCors.ExposeHeaders, profileCors.ExposeHeaders),
		util.GetStringDefault(defaultCors.MaxAge, profileCors.MaxAge),
	}
}

func InitCors(cors *Cors) gin.HandlerFunc {
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
