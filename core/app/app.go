package app

import (
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/env"
	"log"
	"time"
)

func Run(route func(r gin.IRouter)) {
	s := time.Now()
	// load configuration
	conf := env.Load()
	// init log
	env.InitLog()
	// init database
	env.InitDatabase(conf)
	// init server
	env.InitServer(conf, route)
	log.Printf("Started Application in %v seconds", time.Now().Sub(s)/1000)
}
