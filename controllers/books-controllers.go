package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mongo_go/pkg/common/models"
	"net/http"
	"path/filepath"
	"time"
)

//---------------- LIST CATEGORIES --------------------------

func (h handler) ListBooks(ctx *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}

//=========================================================

//---------------- CREATE CATEGORY ------------------------

type CreateBookRequestBody struct {
	Title       string    `json:"title" form:"title"`
	Author      string    `json:"author" form:"author"`
	Description string    `json:"description" form:"description"`
	Price       float64   `json:"price" form:"price"`
	PubDate     time.Time `json:"pubDate" form:"pubDate"`
	CategoryID  uint      `json:"categoryID" form:"categoryID"`
}

func (h handler) CreateBook(ctx *gin.Context) {
	body := CreateBookRequestBody{}

	file, err := ctx.FormFile("image")

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(file.Filename))

	path := "images/" + fileName

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description
	book.Price = body.Price
	book.PubDate = body.PubDate
	book.CategoryID = body.CategoryID
	book.Image = path

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)

}

//=========================================================

//---------------- RETRIEVE BOOK --------------------------

func (h handler) RetrieveBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

//=========================================================

//---------------- UPDATE BOOK ----------------------------

type UpdateBookRequestBody struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	PubDate     time.Time `json:"pubDate"`
	CategoryID  uint      `json:"categoryID"`
}

func (h handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateBookRequestBody{}

	file, err := ctx.FormFile("image")

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(file.Filename))

	path := "images/" + fileName

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description
	book.Price = body.Price
	book.PubDate = body.PubDate
	book.CategoryID = body.CategoryID
	book.Image = path

	h.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}

//=========================================================

//---------------- DELETE BOOK ----------------------------

func (h handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	ctx.Status(http.StatusNoContent)
}

//=========================================================
