package mapper

import (
	"github.com/iDukeLu/believer/core/page"
	"github.com/iDukeLu/believer/core/util"
	"github.com/jinzhu/gorm"
)

var DBS = make(map[string]map[string]*gorm.DB)

func Page(db *gorm.DB, curr int, size int, out interface{}) *gorm.DB {
	curr = util.GetIntDefault(curr, 1)
	size = util.GetIntDefault(size, 10)
	db.Limit(size).Offset(curr - 1).Find(out)
	return db
}

func Count(db *gorm.DB, value interface{}, out interface{}) *gorm.DB {
	db.Model(value).Count(out)
	return db
}

func PageModel(db *gorm.DB, curr int, size int, out interface{}) *page.Page {
	var total int
	Count(db, out, &total)
	Page(db, curr, size, out)
	return &page.Page{Curr: curr, Size: size, Total: total, Records: out}
}
