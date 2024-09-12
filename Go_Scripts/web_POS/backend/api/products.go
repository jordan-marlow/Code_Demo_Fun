package api

import (
	"POS/db"
	"POS/models"
	"net/http"
	"reflect"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func HandleProduct(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		GetProduct(c)
	case http.MethodPut:
		UpdateProduct(c)
	case http.MethodDelete:
		DeleteProduct(c)
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}

func HandleProducts(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		GetProducts(c)
	case http.MethodPost:
		BulkCreateProducts(c)
	case http.MethodPut:
		BulkUpdateProducts(c)
	case http.MethodDelete:
		BulkDeleteProducts(c)
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}

func GetProduct(c *gin.Context) {
	db := db.GetDB()

	id := c.Param("id")

	// Find the existing product by ID
	var product models.Product
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Return the updated product
	c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
	db := db.GetDB()
	var products []models.Product

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Sorting parameters
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc")

	// Search parameter
	search := c.DefaultQuery("search", "")

	query := db.Model(&models.Product{})

	// Apply search filter if provided
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Execute query with pagination and sorting
	result := query.Order(sort + " " + order).Limit(limit).Offset(offset).Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Return the result in JSON format
	c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
	db := db.GetDB()

	id := c.Param("id")

	// Find the existing product by ID
	var product models.Product
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productValue := reflect.ValueOf(&product).Elem()
	updatedProductValue := reflect.ValueOf(updatedProduct)

	for i := 0; i < updatedProductValue.NumField(); i++ {
		field := updatedProductValue.Field(i)
		fieldName := updatedProductValue.Type().Field(i).Name

		// Check if the field can be set and is not zero (has a value)
		if field.IsValid() && field.CanSet() && !field.IsZero() {
			productField := productValue.FieldByName(fieldName)

			// Set the value of the existing product's field
			if productField.IsValid() && productField.CanSet() {
				productField.Set(field)
			}
		}
	}

	// Save the updated product to the database
	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Return the updated product
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	db := db.GetDB()

	id := c.Param("id")

	var product models.Product
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func BulkCreateProducts(c *gin.Context) {
	db := db.GetDB()

	var products []models.Product
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Run the create operation in a transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&products).Error; err != nil {
			// Return an error to rollback the transaction
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Products created successfully", "products": products})
}

func BulkUpdateProducts(c *gin.Context) {
	db := db.GetDB()

	var products []models.Product
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Run the update operation in a transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, product := range products {
			if err := tx.Model(&models.Product{}).Where("id = ?", product.ID).Updates(product).Error; err != nil {
				// Return an error to rollback the transaction
				return err
			}
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Products updated successfully"})
}

func BulkDeleteProducts(c *gin.Context) {
	db := db.GetDB()

	var ids []string
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Run the delete operation in a transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Product{}, ids).Error; err != nil {
			// Return an error to rollback the transaction
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Products deleted successfully"})
}
