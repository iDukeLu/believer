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
}

func getMergeServer(defaultConf *Conf, profileConf *Conf) Server {
	defaultServer := defaultConf.Server
	profileServer := profileConf.Server

	return Server{
		util.GetIntDefault(defaultServer.Port, profileServer.Port),
		util.GetStringDefault(defaultServer.ContextPath, profileServer.ContextPath),
	}
}

func InitServer(c *Conf, route func(r gin.IRouter)) {
	if port := c.Server.Port; port > 0 {
		r := gin.Default()
		if contextPath := c.Server.ContextPath; contextPath != "" {
			route(r.Group(contextPath))
		} else {
			route(r)
		}
		util.LogPanic(r.Run(":" + strconv.Itoa(port)))
	}
	util.LogPanic(errors.New("please use 'server.port' to configure the server port"))
}
