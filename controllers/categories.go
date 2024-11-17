package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kandlagifari/go-books-apps/database"
	"github.com/kandlagifari/go-books-apps/models"
)

func GetCategories(c *gin.Context) {
	rows, err := database.DbConnection.Query(`
        SELECT id, name, created_at, created_by, modified_at, modified_by 
        FROM categories
    `)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.CreatedBy,
			&category.ModifiedAt,
			&category.ModifiedBy,
		); err != nil {
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

	createdBy, _ := c.Get("user")

	query := `
		INSERT INTO categories (name, created_by, created_at)
		VALUES ($1, $2, $3)
	`
	createdAt := time.Now()

	_, err := database.DbConnection.Exec(query, category.Name, createdBy, createdAt)
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

	response := gin.H{
		"id":          category.ID,
		"name":        category.Name,
		"created_at":  category.CreatedAt,
		"created_by":  category.CreatedBy.String,
		"modified_at": category.ModifiedAt,
		"modified_by": category.ModifiedBy.String,
	}

	c.JSON(http.StatusOK, response)
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

func GetBooksByCategoryID(c *gin.Context) {
	id := c.Param("id")

	rows, err := database.DbConnection.Query("SELECT * FROM books WHERE category_id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse book"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedBy, _ := c.Get("user")

	var existingCategory models.Category
	err := database.DbConnection.QueryRow("SELECT * FROM categories WHERE id=$1", id).
		Scan(&existingCategory.ID, &existingCategory.Name, &existingCategory.CreatedAt, &existingCategory.CreatedBy, &existingCategory.ModifiedAt, &existingCategory.ModifiedBy)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	updatedAt := time.Now()

	query := `UPDATE categories SET name=$1, modified_at=$2, modified_by=$3 WHERE id=$4`
	_, err = database.DbConnection.Exec(query, category.Name, updatedAt, updatedBy, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}
