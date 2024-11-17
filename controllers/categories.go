package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kandlagifari/go-books-apps/database"
	"github.com/kandlagifari/go-books-apps/models"
)

func GetCategories(c *gin.Context) {
	rows, err := database.DbConnection.Query("SELECT * FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category"})
			return
		}
		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		INSERT INTO categories (name, created_by)
		VALUES ($1, $2)
	`
	_, err := database.DbConnection.Exec(query, category.Name, "system")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	err := database.DbConnection.QueryRow("SELECT * FROM categories WHERE id=$1", id).
		Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	result, err := database.DbConnection.Exec("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// // TODO: Implement After Books API
// func GetBooksByCategoryID(c *gin.Context) {

// }
