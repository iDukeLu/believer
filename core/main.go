package main

import (
	"github.com/iDukeLu/believer/core/env"
	"github.com/jinzhu/gorm"
)

func main() {

}

var DBS = make(map[string]map[string]*gorm.DB)

type Session struct {
	datasource string
	database   string
}

func DS(datasource string) *Session {
	s := new(Session)
	s.datasource = datasource
	return s
}

func DB(database string) *Session {
	s := new(Session)
	s.database = database
	return s
}

func And(statement string, value interface{}) {

}

func AndIf(condition bool, statement string, value interface{}) {

}

func One() {
	DBS["ADB"]["yx_voyage"].Where()
}
