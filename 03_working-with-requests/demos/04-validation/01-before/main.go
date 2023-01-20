package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TimeoffRequest struct {
	Date   time.Time `json:"date" form:"date" time_format:"2006-01-02" binding:"-"`
	Amount float64   `json:"amount" form:"amount" binding:"-"`
}

func main() {
	router := gin.Default()

	apiGroup := router.Group("/api")

	/*
		Works with message:
		{
			"date": "2022-12-31T00:00:00Z",
			"amount": 8
		}
	*/
	apiGroup.POST("/timeoff", func(c *gin.Context) {
		var timeoffRequest TimeoffRequest
		if err := c.ShouldBindJSON(&timeoffRequest); err == nil {
			c.JSON(http.StatusOK, timeoffRequest)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})

	log.Fatal(router.Run(":3000"))
}
