package env

import (
	"github.com/iDukeLu/believer/core/util"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

const (
	DEFAULT_CONF_PATH = "/app/conf/app.yml"
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
	Url      string `yaml:"url"`
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
	mergeConf.Port = getIntProperty(defaultConf.Port != 0, defaultConf.Port, profileConf.Port)
	mergeConf.ContextPath = getStringProperty(defaultConf.ContextPath != "", defaultConf.ContextPath, profileConf.ContextPath)
	mergeConf.Url = getStringProperty(defaultConf.Url != "", defaultConf.Url, profileConf.Url)
	mergeConf.Username = getStringProperty(defaultConf.Username != "", defaultConf.Username, profileConf.Username)
	mergeConf.Password = getStringProperty(defaultConf.Password != "", defaultConf.Password, profileConf.Password)
	return mergeConf
}

// parse yml []byte to specified struct
func parse(yml []byte) *Conf {
	c := new(Conf)
	e := yaml.Unmarshal(yml, c)
	util.Panic(e)
	return c
}

// read the configuration file of the specified profile
func read(profile string) []byte {
	yml, e := ioutil.ReadFile(getConfFilePath(profile))
	util.Panic(e)
	return yml
}

// get the configuration file path of the specified profile
func getConfFilePath(profile string) string {
	pwd, e := os.Getwd()
	util.Panic(e)

	file := pwd + DEFAULT_CONF_PATH
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
