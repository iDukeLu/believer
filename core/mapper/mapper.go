package mapper

import "github.com/jinzhu/gorm"

var DBS = make(map[string]*gorm.DB)
