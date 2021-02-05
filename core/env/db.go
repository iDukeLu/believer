package env

import (
	"fmt"
	"github.com/iDukeLu/believer/core/mapper"
	"github.com/iDukeLu/believer/core/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strings"
)

type Datasource struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func getMergeDatasource(defaultConf *Conf, profileConf *Conf) []Datasource {
	defaultDatasource := defaultConf.Datasource
	profileDatasource := profileConf.Datasource

	if defaultDatasource != nil && len(defaultDatasource) != 0 {
		return defaultDatasource
	} else {
		return profileDatasource
	}
}

func InitDatabase(c *Conf) {
	log.Println("Initializing Datasource...")
	datasources := c.Datasource
	for _, datasource := range datasources {
		name := strings.Trim(datasource.Name, " ")
		host := strings.Trim(datasource.Host, " ")
		port := datasource.Port
		databases := strings.Trim(datasource.Database, " ")
		username := strings.Trim(datasource.Username, " ")
		password := strings.Trim(datasource.Password, " ")

		if name == "" || host == "" || port <= 0 || databases == "" || username == "" || password == "" {
			return
		}

		for _, database := range strings.Split(databases, ",") {
			database = strings.Trim(database, " ")
			//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
			if connection, e := gorm.Open("mysql", dsn); connection != nil {
				util.LogPanic(e)
				connection.SingularTable(true)
				mapper.DBS[name][database] = connection
			}
		}
		log.Printf("Datasource complete initialization: %v - %v \n", name, getKeys(mapper.DBS[name]))
	}
}

func getKeys(m map[string]*gorm.DB) string {
	var keys string
	for k := range m {
		keys += k + " "
	}
	return keys
}
