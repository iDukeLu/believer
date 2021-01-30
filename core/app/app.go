package app

import (
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/env"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Run(route func(r gin.IRouter)) {
	// load configuration
	conf := env.Load()
	// init database
	env.InitDatabase(conf)
	// init server
	env.InitServer(conf, route)
}
