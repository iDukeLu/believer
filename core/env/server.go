package env

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/util"
	"log"
	"strconv"
	"time"
)

type Server struct {
	Port        int    `yaml:"port"`
	ContextPath string `yaml:"context-path"`
	Cors        Cors   `yaml:"cors"`
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

func InitServer(c *Conf, route func(r gin.IRouter), t time.Time) {
	if port := c.Server.Port; port > 0 {
		log.Println("Initializing Server...")
		mode(c)
		e := gin.Default()
		r := getRouter(c, e)
		route(r)
		middleware(c, r)
		log.Printf("Server started on port(s): %v (http) with context path '%v'", c.Server.Port, c.Server.ContextPath)
		log.Printf("Started Application in %v seconds", time.Now().Sub(t).Milliseconds()/1000)
		util.LogPanic(e.Run(":" + strconv.Itoa(port)))
	}
	util.LogPanic(errors.New("please use 'server.port' to configure the server port"))
}

func getRouter(c *Conf, e *gin.Engine) gin.IRouter {
	if contextPath := c.Server.ContextPath; contextPath != "" {
		return e.Group(contextPath)
	}
	return e
}

func middleware(c *Conf, r gin.IRouter) {
	if c.Server.Cors.Enable {
		r.Use(InitCors(&c.Server.Cors))
		log.Printf("Cors support")
	}
}

func mode(c *Conf) {
	if "prod" == c.Profile {
		gin.SetMode(gin.ReleaseMode)
	}
}
