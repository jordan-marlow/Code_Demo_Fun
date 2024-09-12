package main

import (
	"POS/api"
	"POS/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.GetDB()
	router := gin.Default()
	router.Any("/api/product/:id", api.HandleProduct)
	router.Any("/api/product/:id/", api.HandleProduct)
	router.Any("/api/product/", api.HandleProduct)
	router.Any("/api/products", api.HandleProducts)
	router.Any("/api/products/", api.HandleProducts)

	router.Any("/api/orders", api.HandleOrders)

	// Start the server on port 8080
	router.Run(":8080")

}
