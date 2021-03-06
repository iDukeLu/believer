package app

import (
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/env"
	"time"
)

func Run(route func(r gin.IRouter)) {
	s := time.Now()

	// init log
	env.InitLog()

	// init banner
	env.InitBanner()

	// load configuration
	conf := env.Load()

	// init database
	env.InitDatabase(conf)

	// init server
	env.InitServer(conf, route, s)

}
