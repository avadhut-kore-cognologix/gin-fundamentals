package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}
	r := gin.Default()
	registerRoutes(r)

	r.Run()

}

func registerRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	r.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	r.POST("/employee", func(c *gin.Context) {
		c.Request.Method = http.MethodGet
		c.Redirect(http.StatusFound, "/employee")
	})

	r.Static("/public", "./public")
}
