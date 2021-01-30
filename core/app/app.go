package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iDukeLu/believer/core/env"
	"github.com/iDukeLu/believer/core/util"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func Run(route func(r gin.IRouter)) {
	// load configuration
	conf := env.Load()
	// init server
	initServer(conf, route)
	// init database
	initDatabase(conf)
}

func initServer(c *env.Conf, route func(r gin.IRouter)) {
	if port := c.Server.Port; port > 0 {
		r := gin.Default()
		if contextPath := c.Server.ContextPath; contextPath != "" {
			route(r.Group(contextPath))
		} else {
			route(r)
		}
		util.Panic(r.Run(":" + string(port)))
	}
	util.Panic(errors.New("please use 'server.port' to configure the server port"))
}

func initDatabase(c *env.Conf) {
	datasource := c.Datasource
	host := datasource.Host
	port := datasource.Port
	database := datasource.Database
	username := datasource.Username
	password := datasource.Password

	if host == "" || port == 0 || database == "" || username == "" || password == "" {
		return
	}

	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	db, e := gorm.Open("mysql", dsn)
	util.Panic(e)
	DB = db
}
