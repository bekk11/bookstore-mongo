package controllers

import (
	"github.com/gin-gonic/gin"
	"mongo_go/pkg/common/models"
	"net/http"
)

//---------------- LIST CATEGORIES --------------------------

func (h handler) ListCategories(ctx *gin.Context) {
	var categories []models.Category

	if result := h.DB.Find(&categories); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
	}

	ctx.JSON(http.StatusOK, &categories)
}

//=========================================================

//---------------- CREATE CATEGORY ------------------------

type CreateCategoryRequestBody struct {
	Name string `json:"name"`
}

func (h handler) CreateCategory(ctx *gin.Context) {
	body := CreateCategoryRequestBody{}

	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var category models.Category

	category.Name = body.Name

	if result := h.DB.Create(&category); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &category)
}

//=========================================================

//---------------- RETRIEVE CATEGORY ----------------------

func (h handler) RetrieveCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	var category models.Category

	if result := h.DB.First(&category, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &category)
}

//=========================================================

//---------------- UPDATE CATEGORY ------------------------

type UpdateCategoryRequestBody struct {
	Name string `json:"name"`
}

func (h handler) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateCategoryRequestBody{}

	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var category models.Category

	if result := h.DB.First(&category, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	category.Name = body.Name

	h.DB.Save(&category)

	ctx.JSON(http.StatusOK, &category)
}

//=========================================================

//---------------- DELETE CATEGORY ------------------------

func (h handler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	var category models.Category

	if result := h.DB.First(&category, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&category)

	ctx.Status(http.StatusNoContent)
}

//=========================================================
