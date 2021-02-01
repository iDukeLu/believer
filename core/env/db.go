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
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Databases string `yaml:"databases"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

func getMergeDatasource(defaultConf *Conf, profileConf *Conf) Datasource {
	defaultDatasource := defaultConf.Datasource
	profileDatasource := profileConf.Datasource

	return Datasource{
		util.GetStringDefault(defaultDatasource.Host, profileDatasource.Host),
		util.GetIntDefault(defaultDatasource.Port, profileDatasource.Port),
		util.GetStringDefault(defaultDatasource.Databases, profileDatasource.Databases),
		util.GetStringDefault(defaultDatasource.Username, profileDatasource.Username),
		util.GetStringDefault(defaultDatasource.Password, profileDatasource.Password),
	}
}

func InitDatabase(c *Conf) {
	log.Println("Initializing Datasource...")
	datasource := c.Datasource
	host := datasource.Host
	port := datasource.Port
	databases := datasource.Databases
	username := datasource.Username
	password := datasource.Password

	if host == "" || port == 0 || databases == "" || username == "" || password == "" {
		return
	}

	for _, database := range strings.Split(databases, ",") {
		database = strings.Trim(database, " ")
		//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
		if connection, e := gorm.Open("mysql", dsn); connection != nil {
			util.LogPanic(e)
			connection.SingularTable(true)
			mapper.DBS[database] = connection
		}
	}
	log.Printf("Datasource complete initialization: %v \n", getKeys(&mapper.DBS))
}

func getKeys(m *map[string]*gorm.DB) []string {
	keys := make([]string, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}
	return keys
}
