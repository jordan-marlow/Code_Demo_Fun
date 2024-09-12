package main

import (
	"library/api"
	"library/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.GetDB()
	router := gin.Default()
	router.Any("/api/author/:id", api.HandleAuthor)
	router.Any("/api/author/:id/", api.HandleAuthor)
	router.Any("/api/author/", api.HandleAuthor)
	router.Any("/api/authors", api.HandleAuthors)
	router.Any("/api/authors/", api.HandleAuthors)

	// Start the server on port 8080
	router.Run(":8080")

}
