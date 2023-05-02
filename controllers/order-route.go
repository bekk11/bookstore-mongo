package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/orders")

	routes.GET("/", h.ListOrders)
	routes.POST("/", h.CreateOrder)

	routes.GET("/:id", h.RetrieveOrder)
	routes.PUT("/:id", h.UpdateOrder)
	routes.DELETE("/:id", h.DestroyOrder)
}
