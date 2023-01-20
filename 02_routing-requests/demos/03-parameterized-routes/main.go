package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

)

func main() {
	router := gin.Default()

	router.GET("/employees/:username/*rest", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"username": c.Param("username"),
			"rest":     c.Param("rest"),
		})
	})

	log.Fatal(router.Run(":3000"))
}
