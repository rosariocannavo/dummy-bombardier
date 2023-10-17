package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyData struct {
	Value string `json:"value"`
}

func main() {
		r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, GET request!",
		})
	})


	r.POST("/data", func(c *gin.Context) {
		var inputData MyData
		if err := c.ShouldBindJSON(&inputData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		message := "Data received via POST request: " + inputData.Value

		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})


	// Run the server on port 8080
	r.Run(":8080")
}
