package main

import (
	"addressapi/db"
	"addressapi/models"
	// "gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"addressapi/handlers"
	"addressapi/auth"
)

func main() {
	if err := db.Connect();err != nil {
		panic("Failed to connect to the database")
	}
	db.DB.AutoMigrate(&models.Customer{})

	r := gin.Default()

	r.GET("/customers", auth.Auth(),handlers.GetCustomers)
	r.GET("/customers/:id", auth.Auth(),handlers.GetByid)
	r.POST("/customers", auth.Auth(),handlers.PostCustomer)
	r.PUT("/customers/:id", auth.Auth(),handlers.UpdateCustomer)
	r.DELETE("/customers/:id", auth.Auth(),handlers.DeleteCustomer)

	r.Run(":8080")
}
