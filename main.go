package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	
	r.POST("/customer", func(c *gin.Context) {
		// Handle customer creation
		c.JSON(201, gin.H{"message": "Customer created"})
	})
	r.GET("/customer/:customerId", func(c *gin.Context) {
	

	}	)
	r.PUT("/customer/customerId", func(c *gin.Context) {			
	
	})
	r.DELETE("/customer/customerId", func(c *gin.Context) {	
	
	})	

	r.Run(":8080") // Start the server on port 8080
}