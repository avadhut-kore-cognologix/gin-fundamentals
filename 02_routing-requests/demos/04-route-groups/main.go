package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

)

func main() {
	router := gin.Default()

	adminGroup := router.Group("/admin")
	adminGroup.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "Page to adminster users")
	})
	adminGroup.GET("/roles", func(c *gin.Context) {
		c.String(http.StatusOK, "Page to adminster roles")
	})
	adminGroup.GET("/policies", func(c *gin.Context) {
		c.String(http.StatusOK, "Page to adminster vacation policies")
	})

	log.Fatal(router.Run(":3000"))
}
