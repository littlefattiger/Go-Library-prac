package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":    "apple",
			"message": "hello world!",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		// 方法二：使用结构体
		var msg struct {
			Name    string `json:"name"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	r.Run(":9090")
}
