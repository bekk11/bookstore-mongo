package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/book")

	routes.GET("/", h.ListBooks)
	routes.POST("/", h.CreateBook)

	routes.GET("/:id", h.RetrieveBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
