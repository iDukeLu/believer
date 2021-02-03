package req

import (
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
)

func Params(c *gin.Context, obj interface{}) interface{} {
	if e := c.BindQuery(obj); e != nil {
		log.Panic(e)
	}
	return obj
}

func Body(c *gin.Context, obj interface{}) interface{} {
	if e := c.Bind(obj); e != nil {
		log.Panic(e)
	}
	return obj
}

func Headers(c *gin.Context, obj interface{}) interface{} {
	if e := c.BindHeader(obj); e != nil {
		log.Panic(e)
	}
	return obj
}

func File(c *gin.Context, name string) *multipart.FileHeader {
	f, e := c.FormFile(name)
	if e != nil {
		log.Panic(e)
	}
	return f
}

func Xml(c *gin.Context, obj interface{}) interface{} {
	if e := c.BindXML(obj); e != nil {
		log.Panic(e)
	}
	return obj
}
