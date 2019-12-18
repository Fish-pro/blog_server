package main

import (
	"blog_server/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// 命令行参数
	// 从配置读取参数
	var err error
	conf := config.LoadConf()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
