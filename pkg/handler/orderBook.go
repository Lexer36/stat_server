package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"statServer/pkg/models"
)

func (h *Handler) getOrderBook(c *gin.Context) {
	var input models.Identity
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	res, err := h.services.OrderBook.GetOrderBook(input.ExchangeName, input.Pair)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) saveOrderBook(c *gin.Context) {
	var input models.OrderBook
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	err := h.services.OrderBook.SaveOrderBook(input.ExchangeName, input.Pair, input.Asks, input.Bids)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
