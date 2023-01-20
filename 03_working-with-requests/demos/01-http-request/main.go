package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/*rest", func(c *gin.Context) {
		url := c.Request.URL.String()
		headers := c.Request.Header
		cookies := c.Request.Cookies()

		c.IndentedJSON(http.StatusOK, gin.H{
			"url":     url,
			"headers": headers,
			"cookies": cookies,
		})
	})

	log.Fatal(router.Run(":3000"))
}
