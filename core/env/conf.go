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

type conf struct {
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

func Load() *conf {
	//应该是 绝对地址
	pwd, e := os.Getwd()
	util.Panic(e)

	yamlFile, e := ioutil.ReadFile(pwd + DEFAULT_CONF_PATH)
	util.Panic(e)

	c := new(conf)
	e = yaml.Unmarshal(yamlFile, c)
	util.Panic(e)

	return c
}
