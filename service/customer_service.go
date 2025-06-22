package service

import (
	"customer-service/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct{
	ID int `json:"id"`
	Name string `json:"name,omitempty"`
	Email string `json:"email"`
	Address string `json:"address,omitempty"`
}


func createCustomer(db *db.PostgresDB, c *gin.Context) (int, *Customer, error) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if len(customer.Email) == 0 {
		return http.StatusBadRequest, nil, fmt.Errorf("email cannot be empty")
	}

    stmt := `INSERT INTO customers (name, email, address) VALUES ($1, $2, $3) RETURNING id`
	err := db.DB.QueryRow(stmt, customer.Name, customer.Email, customer.Address).Scan(&customer.ID)
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("failed to create customer: %w", err)
	}

	return http.StatusCreated, &customer, nil
}

func getCustomer()

func updateCustomer()

func deleteCustomer() {}
