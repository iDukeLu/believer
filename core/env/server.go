package env

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/util"
	"strconv"
)

type Server struct {
	Port        int    `yaml:"port"`
	ContextPath string `yaml:"context-path"`
	Cors        Cors   `yaml:"InitCors"`
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

func InitServer(c *Conf, route func(r gin.IRouter)) {
	if port := c.Server.Port; port > 0 {
		e := gin.Default()
		r := getRouter(c, e)
		route(r)
		middleware(c, r)
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

func middleware(c *Conf, r gin.IRouter) {
	if c.Server.Cors.Enable {
		r.Use(InitCors(&c.Server.Cors))
	}
}
