package service

import (
	"statServer/pkg/models"
	"statServer/pkg/repository"
)

type OrderBookService struct {
	repo repository.OrderBook
}

func NewOrderBookService(repo repository.OrderBook) *OrderBookService {
	return &OrderBookService{repo: repo}
}

func (s *OrderBookService) GetOrderBook(exchangeName, pair string) (*models.OrderBook, error) {
	// Передача входных параметров репозиторию для получения данных
	orderBook, err := s.repo.GetOrderBook(exchangeName, pair)
	if err != nil {
		return nil, err
	}

	return orderBook, nil
}

func (s *OrderBookService) SaveOrderBook(exchangeName, pair string, asks, bids []*models.DepthOrder) error {
	// Передача входных параметров репозиторию для сохранения данных
	err := s.repo.SaveOrderBook(exchangeName, pair, asks, bids)
	if err != nil {
		return err
	}

	return nil
}
