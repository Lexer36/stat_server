package handler

import (
	"github.com/gin-gonic/gin"
	"statServer/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	orderBook := router.Group("/order-book")
	{
		orderBook.GET("/get", h.getOrderBook)
		orderBook.POST("/save", h.saveOrderBook)
	}

	orderHistory := router.Group("/order-history")
	{
		orderHistory.POST("/get", h.getOrderHistory)
		orderHistory.POST("/save", h.saveOrder)
	}

	return router
}
