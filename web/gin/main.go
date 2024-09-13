package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 官方文档 https://gin-gonic.com/zh-cn/docs/
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(":8888"); err != nil {
		panic(err)
	}
}
