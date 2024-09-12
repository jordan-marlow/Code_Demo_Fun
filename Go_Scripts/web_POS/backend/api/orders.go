package api

import (
	"net/http"

	"POS/db"
	"POS/models"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleOrders(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		BulkCreateOrders(c)
	case http.MethodGet:
		GetOrders(c)
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}

func GetOrder(c *gin.Context) {
	db := db.GetDB()

	id := c.Param("id")

	// Find the existing product by ID
	var order models.Order
	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Return the updated product
	c.JSON(http.StatusOK, order)

}

// GetOrders retrieves orders with optional pagination, sorting, and filtering.
func GetOrders(c *gin.Context) {
	db := db.GetDB()
	var orders []models.Order

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Sorting parameters
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc")

	// Search parameter
	search := c.DefaultQuery("search", "")

	// Base query
	query := db.Preload("Products").Model(&models.Order{})

	// Apply search filter if provided
	if search != "" {
		query = query.Where("customer LIKE ?", "%"+search+"%")
	}

	// Execute query with pagination and sorting
	result := query.Order(sort + " " + order).Limit(limit).Offset(offset).Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Return the result in JSON format
	c.JSON(http.StatusOK, orders)
}

func BulkCreateOrders(c *gin.Context) {
	db := db.GetDB()

	var orders []models.Order
	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Iterate over each order to calculate the total price
	for i := range orders {
		var totalCents int64
		var orderProducts []models.Product
		for _, product := range orders[i].Products {

			var fetchedProduct models.Product
			if err := db.First(&fetchedProduct, product.ID).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
				return
			}

			// Calculate total price
			totalCents += fetchedProduct.PriceCents
			orderProducts = append(orderProducts, fetchedProduct)
		}

		orders[i].TotalCents = totalCents
		orders[i].Products = orderProducts
	}

	// Perform bulk create in a transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&orders).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orders created successfully", "orders": orders})
}

// BulkUpdateOrders updates multiple orders in bulk.
func BulkUpdateOrders(c *gin.Context) {
	db := db.GetDB()

	var orders []models.Order
	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		for _, order := range orders {
			// Update each order individually
			if err := tx.Model(&order).Association("Products").Replace(order.Products); err != nil {
				return err
			}
			if err := tx.Save(&order).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orders updated successfully"})
}

// BulkDeleteOrders deletes multiple orders by their IDs.
func BulkDeleteOrders(c *gin.Context) {
	db := db.GetDB()

	// Expecting a JSON array of order IDs to delete
	var orderIDs []uint
	if err := c.ShouldBindJSON(&orderIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		for _, id := range orderIDs {
			if err := tx.Delete(&models.Order{}, id).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orders deleted successfully"})
}
