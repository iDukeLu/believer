package env

import (
	"github.com/iDukeLu/believer/core/util"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultConfPath        = "/app/conf/app"
	defaultConfFilesSuffix = ".yml"
)

type Conf struct {
	Profile    string     `yaml:"profile"`
	Server     Server     `yaml:"server"`
	Datasource Datasource `yaml:"datasource"`
}

// load、parse、merge configuration file
func Load() *Conf {
	defaultConf := parse(read(""))
	log.Printf("The following profiles are active: %v", getActivityProfile(defaultConf))
	if profile := defaultConf.Profile; profile != "" {
		profileConf := parse(read(profile))
		return merge(defaultConf, profileConf)
	}
	log.Println("Configuration finished loading")
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

// get the configuration file path of the specified profile
func getConfFilePath(profile string) string {
	pwd, e := os.Getwd()
	util.LogPanic(e)

	path := pwd + defaultConfPath
	if profile != "" {
		path += "-" + profile
	}
	return path + defaultConfFilesSuffix
}

func getActivityProfile(c *Conf) string {
	if p := c.Profile; p != "" {
		return p
	}
	return "default"
}
