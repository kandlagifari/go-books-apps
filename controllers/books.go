package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kandlagifari/go-books-apps/database"
	"github.com/kandlagifari/go-books-apps/models"
)

func GetBooks(c *gin.Context) {
	rows, err := database.DbConnection.Query("SELECT * FROM books")
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

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	query := `
		INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := database.DbConnection.Exec(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, "system")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	err := database.DbConnection.QueryRow("SELECT * FROM books WHERE id=$1", id).
		Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	result, err := database.DbConnection.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
