package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type TimeoffRequest struct {
	Date   time.Time `json:"date" form:"date" time_format:"2006-01-02" binding:"required,future"`
	Amount float64   `json:"amount" form:"amount" binding:"required,gt=0"`
}

var ValidatorFuture validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		return date.After(time.Now())
	}
	return true
}

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("future", ValidatorFuture)
	}

	apiGroup := router.Group("/api")

	/*
		Works with message:
		{
			"date": "2022-12-31T00:00:00Z",
			"amount": 8
		}

		Fails with messages:
		{
			"date": "2022-12-31T00:00:00Z",
			"amount": -4											// invalid amount
		}

		{
			"date": "2000-01-01T00:00:00Z",  	// invalid date
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
