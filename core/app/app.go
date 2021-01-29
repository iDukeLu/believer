package app

import (
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/env"
	"github.com/iDukeLu/believer/core/util"
	"net/http"
)

func Run() {
	// load configuration
	conf := env.Load()
	// init server
	initServer(conf)
	// init database
}

func initServer(c *env.Conf) {
	r := gin.Default()
	contextPath := r.Group(c.ContextPath)
	util.Panic(r.Run(":" + string(c.Port)))
}

func test(c *gin.Context) {
	c.String(http.StatusOK, "v1 login")
}
