package service

import (
	"statServer/pkg/models"
	"statServer/pkg/repository"
)

type OrderBook interface {
	GetOrderBook(exchangeName, pair string) (*models.OrderBook, error)
	SaveOrderBook(exchangeName, pair string, asks, bids []*models.DepthOrder) error
}

type OrderHistory interface {
	GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error)
	SaveOrder(client *models.Client, order *models.HistoryOrder) error
}

type Service struct {
	OrderBook
	OrderHistory
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		OrderBook:    NewOrderBookService(repos.OrderBook),
		OrderHistory: NewOrderHistoryService(repos.OrderHistory),
	}
}
