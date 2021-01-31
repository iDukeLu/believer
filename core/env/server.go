package env

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/middleware"
	"github.com/iDukeLu/believer/core/util"
	"strconv"
)

type Server struct {
	Port        int    `yaml:"port"`
	ContextPath string `yaml:"context-path"`
	Cors        Cors   `yaml:"cors"`
}

type Cors struct {
	Enable           bool   `yaml:"enable"`
	AllowOrigin      string `yaml:"allow-origin"`
	AllowMethods     string `yaml:"allow-methods"`
	AllowHeaders     string `yaml:"allow-headers"`
	AllowCredentials string `yaml:"allow-credentials"`
	ExposeHeaders    string `yaml:"expose-headers"`
	MaxAge           string `yaml:"max-age"`
}

func getMergeServer(defaultConf *Conf, profileConf *Conf) Server {
	defaultServer := defaultConf.Server
	profileServer := profileConf.Server

	return Server{
		util.GetIntDefault(defaultServer.Port, profileServer.Port),
		util.GetStringDefault(defaultServer.ContextPath, profileServer.ContextPath),
		getMergeCors(&defaultServer.Cors, &profileServer.Cors),
	}
}

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

func InitServer(c *Conf, route func(r gin.IRouter)) {
	if port := c.Server.Port; port > 0 {
		e := gin.Default()
		r := getRouter(c, e)
		route(r)
		middle(c, r)
		util.LogPanic(e.Run(":" + strconv.Itoa(port)))
	}
	util.LogPanic(errors.New("please use 'server.port' to configure the server port"))
}

func getRouter(c *Conf, r *gin.Engine) gin.IRouter {
	if contextPath := c.Server.ContextPath; contextPath != "" {
		return r.Group(contextPath)
	}
	return r
}

func middle(c *Conf, r gin.IRouter) {
	if c.Server.Cors.Enable {
		r.Use(middleware.Cors(&c.Server.Cors))
	}
}
