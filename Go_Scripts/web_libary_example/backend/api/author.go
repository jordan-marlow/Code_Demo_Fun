package api

import (
	"fmt"
	"library/db"
	"library/models"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
)

func HandleAuthors(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		getAuthors(c)
	case http.MethodPut:
		bulkUpdateAuthors(c)
	case http.MethodPost:
		bulkCreateAuthors(c)
	case http.MethodDelete:
		bulkDeleteAuthors(c)
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid Method"})
	}
}

func HandleAuthor(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		getAuthors(c)
	case http.MethodPut:
		updateAuthor(c)
	case http.MethodDelete:
		deleteAuthor(c)
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid Method"})
	}
}

func bulkCreateAuthors(c *gin.Context) {
	db := db.GetDB()

	var authors []models.Author

	if err := c.ShouldBindJSON(&authors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	fmt.Println(authors)
	// Run the create operation in a transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&authors).Error; err != nil {
			// Return an error to rollback the transaction
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create authors"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Authors created successfully", "authors": authors})
}

func bulkUpdateAuthors(c *gin.Context) {
	db := db.GetDB()

	var authors []models.Author

	if err := c.ShouldBindJSON(&authors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		for _, author := range authors {
			// Update the author by ID, only the changed fields
			if err := tx.Model(&models.Author{}).Where("id = ?", author.ID).Updates(author).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update authors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Authors updated successfully", "authors": authors})
}

func bulkDeleteAuthors(c *gin.Context) {
	db := db.GetDB()

	var ids []ulid.ULID
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Run the delete operation in a transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if err := tx.Where("id = ?", id).Delete(&models.Author{}).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete authors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Authors deleted successfully"})

}

func getAuthors(c *gin.Context) {
	db := db.GetDB()
	var products []models.Author

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Sorting parameters
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc")

	// Search parameter
	search := c.DefaultQuery("search", "")

	query := db.Model(&models.Author{})

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

func updateAuthor(c *gin.Context) {
	db := db.GetDB()

	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Author{}).Where("id = ?", author.ID).Updates(author).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author updated successfully", "author": author})
}

func deleteAuthor(c *gin.Context) {
	db := db.GetDB()

	idParam := c.Param("id")
	var author models.Author
	id, err := ulid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not convert id to ULID"})
		return
	}

	if err := db.Where("id = ?", id).First(&author).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	if err := db.Delete(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
