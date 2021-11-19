package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"wechat": "flysnow_org", "blog": "www.flysnow.org"})
	})
	r.POST("/upload", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	r.Run(":8080")
}
