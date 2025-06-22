package service

import (
	"customer-service/db"

	"github.com/gin-gonic/gin"
)

type App struct {
	db *db.PostgresDB
}

func GetApp(db db.PostgresDB) *App {
	return &App{
		db: &db,
	}
}


func (a *App) PostHandler(c *gin.Context) {
	
	status,customer, err := createCustomer(a.db,c)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, customer)
}

func (a *App) GetHandler(c *gin.Context) {
	customerId := c.Param("customerId")
	c.JSON(200, gin.H{"message": "Customer retrieved", "customerId": customerId})
}

func (a *App) PutHandler(c *gin.Context) {
	customerId := c.Param("customerId")
	c.JSON(200, gin.H{"message": "Customer updated", "customerId": customerId})
}

func (a *App) DeleteHandler(c *gin.Context) {
	customerId := c.Param("customerId")
	c.JSON(200, gin.H{"message": "Customer deleted", "customerId": customerId})
}