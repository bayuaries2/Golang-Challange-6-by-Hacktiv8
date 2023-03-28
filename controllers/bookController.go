package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var Books = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.ID = fmt.Sprintf("%d", len(Books)+1)
	Books = append(Books, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBookById(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range Books {
		if bookID == book.ID {
			condition = true
			Books[i] = updateBook
			Books[i].ID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_massage": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Massage": fmt.Sprintf("Book with id %v Updated", bookID),
	})
}

func GetAllBook(ctx *gin.Context) {
	if len(Books) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_massage": "Data tidak tersedia",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Massage": "Data Tersedia",
		"book":    Books,
	})
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false

	var getOneBook Book

	for in, book := range Books {
		if bookID == book.ID {
			condition = true
			getOneBook = Books[in]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_massage": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Massage": fmt.Sprintf("Get book with id %v", bookID),
		"book":    getOneBook,
	})

}

func DeleteBookById(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false

	var getBookIndex int

	for in, book := range Books {
		if bookID == book.ID {
			condition = true
			getBookIndex = in
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_massage": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	copy(Books[getBookIndex:], Books[getBookIndex+1:])
	Books[len(Books)-1] = Book{}
	Books = Books[:len(Books)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"Massage": fmt.Sprintf("Book with id %v has been succsesfully deleted", bookID),
	})
}
