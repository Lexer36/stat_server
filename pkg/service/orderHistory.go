package service

import (
	"statServer/pkg/models"
	"statServer/pkg/repository"
)

type OrderHistoryService struct {
	repo repository.OrderHistory
}

func NewOrderHistoryService(repo repository.OrderHistory) *OrderHistoryService {
	return &OrderHistoryService{repo: repo}
}

func (s *OrderHistoryService) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	res, err := s.repo.GetOrderHistory(client)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *OrderHistoryService) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	err := s.repo.SaveOrder(client, order)
	if err != nil {
		return err
	}
	return nil
}
