package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/category")
	routes.GET("/", h.ListCategories)
	routes.POST("/", h.CreateCategory)
	routes.GET("/:id", h.RetrieveCategory)
	routes.PUT("/:id", h.UpdateCategory)
	routes.DELETE("/:id", h.DeleteCategory)
}
