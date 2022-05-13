package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/xxx", "./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/home.html", "templates/users/home.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "templates/home.html", gin.H{
			"title": "index",
		})
	})
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/home.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/home.html", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	r.Run(":9090")
}
