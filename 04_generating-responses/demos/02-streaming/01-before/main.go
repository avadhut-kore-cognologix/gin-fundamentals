package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("/", "./index.html")

	router.GET("/tale_of_two_cities", func(c *gin.Context) {
		f, err := os.Open("./a_tale_of_two_cities.txt")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer f.Close()
		data, err := io.ReadAll(f)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Data(http.StatusOK, "text/plain", data)
	})

	log.Fatal(router.Run(":3000"))
}
