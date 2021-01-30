package env

import (
	"github.com/iDukeLu/believer/core/util"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

const (
	defaultConfPath = "/app/conf/app.yml"
)

type Conf struct {
	Profile    string `yaml:"profile"`
	Server     `yaml:"server"`
	Datasource `yaml:"datasource"`
}

type Server struct {
	Port        int    `yaml:"port"`
	ContextPath string `yaml:"context-path"`
}

type Datasource struct {
	Host     string `yaml:"url"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// load and parse configuration file
func Load() *Conf {
	return doLoad()
}

func doLoad() *Conf {
	defaultConf := parse(read(""))
	if profile := defaultConf.Profile; profile != "" {
		profileConf := parse(read(profile))
		return merge(defaultConf, profileConf)
	}
	return defaultConf
}

// merge configuration file
func merge(defaultConf *Conf, profileConf *Conf) *Conf {
	mergeConf := new(Conf)
	mergeConf.Profile = defaultConf.Profile
	mergeConf.Server = getMergeServer(defaultConf, profileConf)
	mergeConf.Datasource = getMergeDatasource(defaultConf, profileConf)
	return mergeConf
}

// parse yml []byte to specified struct
func parse(yml []byte) *Conf {
	c := new(Conf)
	e := yaml.Unmarshal(yml, c)
	util.LogPanic(e)
	return c
}

// read the configuration file of the specified profile
func read(profile string) []byte {
	yml, e := ioutil.ReadFile(getConfFilePath(profile))
	util.LogPanic(e)
	return yml
}

func getMergeServer(defaultConf *Conf, profileConf *Conf) Server {
	defaultServer := defaultConf.Server
	profileServer := profileConf.Server
	return Server{
		getIntProperty(defaultServer.Port != 0, defaultServer.Port, profileServer.Port),
		getStringProperty(defaultServer.ContextPath != "", defaultServer.ContextPath, profileServer.ContextPath),
	}
}

func getMergeDatasource(defaultConf *Conf, profileConf *Conf) Datasource {
	defaultDatasource := defaultConf.Datasource
	profileDatasource := profileConf.Datasource
	return Datasource{
		getStringProperty(defaultDatasource.Host != "", defaultDatasource.Host, profileDatasource.Host),
		getIntProperty(defaultDatasource.Port != 0, defaultDatasource.Port, profileDatasource.Port),
		getStringProperty(defaultDatasource.Database != "", defaultDatasource.Database, profileDatasource.Database),
		getStringProperty(defaultDatasource.Username != "", defaultDatasource.Username, profileDatasource.Username),
		getStringProperty(defaultDatasource.Password != "", defaultDatasource.Password, profileDatasource.Password),
	}
}

// get the configuration file path of the specified profile
func getConfFilePath(profile string) string {
	pwd, e := os.Getwd()
	util.LogPanic(e)

	file := pwd + defaultConfPath
	if profile != "" {
		file += "-" + profile
	}
	return file
}

func getStringProperty(con bool, v1 string, v2 string) string {
	if con {
		return v1
	}
	return v2
}

func getIntProperty(con bool, v1 int, v2 int) int {
	if con {
		return v1
	}
	return v2
}
