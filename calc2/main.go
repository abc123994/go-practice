package main

import (
	calc "calc2/router/calc"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello, gin")
	})
	r.LoadHTMLGlob("html/*")
	r.GET("index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/calc", calc.Calc)

	r.Run(":80") // listen and serve on 0.0.0.0:8080
}
