package env

import (
	"fmt"
	"github.com/iDukeLu/believer/core/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Datasource struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func getMergeDatasource(defaultConf *Conf, profileConf *Conf) Datasource {
	defaultDatasource := defaultConf.Datasource
	profileDatasource := profileConf.Datasource

	return Datasource{
		util.GetStringDefault(defaultDatasource.Host, profileDatasource.Host),
		util.GetIntDefault(defaultDatasource.Port, profileDatasource.Port),
		util.GetStringDefault(defaultDatasource.Database, profileDatasource.Database),
		util.GetStringDefault(defaultDatasource.Username, profileDatasource.Username),
		util.GetStringDefault(defaultDatasource.Password, profileDatasource.Password),
	}
}

func InitDatabase(c *Conf) *gorm.DB {
	datasource := c.Datasource
	host := datasource.Host
	port := datasource.Port
	database := datasource.Database
	username := datasource.Username
	password := datasource.Password

	if host != "" && port != 0 && database != "" && username != "" && password != "" {
		//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
		if db, e := gorm.Open("mysql", dsn); db != nil {
			util.LogPanic(e)
			db.SingularTable(true)
			return db
		}
	}
	return nil
}
