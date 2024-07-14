package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"statServer/pkg/models"
)

type InputGetOrderHistory struct {
	Client *models.Client `json:"client"`
}

type InputSaveOrderHistory struct {
	InputGetOrderHistory
	Order *models.HistoryOrder `json:"order"`
}

func (h *Handler) getOrderHistory(c *gin.Context) {
	var input InputGetOrderHistory
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	res, err := h.services.OrderHistory.GetOrderHistory(input.Client)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) saveOrder(c *gin.Context) {
	var input InputSaveOrderHistory
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	err := h.services.OrderHistory.SaveOrder(input.Client, input.Order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
