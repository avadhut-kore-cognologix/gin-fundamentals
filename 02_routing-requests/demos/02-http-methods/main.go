package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

)

func main() {
	router := gin.Default()

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	router.POST("/employee", func(c *gin.Context) {
		c.String(http.StatusOK, "New request POSTed successfully!")
	})

	log.Fatal(router.Run(":3000"))
}
