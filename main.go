package main

import (
	"customer-service/db"
	"customer-service/service"

	 "github.com/gin-gonic/gin"
)

func main(){


	secret := db.GetSecretValue()
	db := db.GetDB(secret)

	a := service.GetApp(*db)

	r := gin.Default()
	
	r.POST("/customer", a.PostHandler)
	r.GET("/customer/:customerId", a.GetHandler)
	r.PUT("/customer/customerId", a.PutHandler)
	r.DELETE("/customer/customerId", a.DeleteHandler)	

	r.Run(":8080") // Start the server on port 8080
}