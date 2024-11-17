package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kandlagifari/go-books-apps/database"
	"github.com/kandlagifari/go-books-apps/models"
	"github.com/lib/pq"
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

	updatedBy, exists := c.Get("user")
	if !exists || updatedBy == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User context missing"})
		return
	}

	var categoryExists bool
	err := database.DbConnection.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id=$1)", book.CategoryID).Scan(&categoryExists)
	if err != nil || !categoryExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id"})
		return
	}

	createdAt := time.Now()

	query := `
		INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err = database.DbConnection.Exec(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, int(book.Price), book.TotalPage, book.Thickness, book.CategoryID, updatedBy, createdAt)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			c.JSON(http.StatusConflict, gin.H{"error": "Book title must be unique"})
			return
		}

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

	response := gin.H{
		"id":           book.ID,
		"title":        book.Title,
		"description":  book.Description,
		"image_url":    book.ImageURL,
		"release_year": book.ReleaseYear,
		"price":        book.Price,
		"total_page":   book.TotalPage,
		"thickness":    book.Thickness,
		"category_id":  book.CategoryID,
		"created_at":   book.CreatedAt,
		"created_by":   book.CreatedBy.String,
		"modified_at":  book.ModifiedAt,
		"modified_by":  book.ModifiedBy.String,
	}

	c.JSON(http.StatusOK, response)
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

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedBy, _ := c.Get("user")

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	var existingBook models.Book
	err := database.DbConnection.QueryRow("SELECT * FROM books WHERE id=$1", id).
		Scan(&existingBook.ID, &existingBook.Title, &existingBook.Description, &existingBook.ImageURL, &existingBook.ReleaseYear, &existingBook.Price, &existingBook.TotalPage, &existingBook.Thickness, &existingBook.CategoryID, &existingBook.CreatedAt, &existingBook.CreatedBy, &existingBook.ModifiedAt, &existingBook.ModifiedBy)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	updatedAt := time.Now()

	query := `
		UPDATE books 
		SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, category_id=$8, modified_at=$9, modified_by=$10 
		WHERE id=$11
	`
	_, err = database.DbConnection.Exec(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, updatedAt, updatedBy, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
