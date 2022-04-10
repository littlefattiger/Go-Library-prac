package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index ...")
	c.JSON(http.StatusOK, gin.H{"msg": "index"})
}
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
	fmt.Println("m1 out..")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in..")
	c.Next()
	fmt.Println("m2 out..")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {

		} else {
			c.Next()
		}
	}
}
func main() {

	r := gin.Default()
	r.Use(m1, m2, authMiddleware(true))
	r.GET("/index", indexHandler)

	r.Run(":9090")
}
