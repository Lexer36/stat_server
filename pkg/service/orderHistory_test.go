package service

import (
	"errors"
	"reflect"
	"statServer/pkg/models"
	"testing"
)

type mockOrderHistory struct{}

func (m *mockOrderHistory) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	return []*models.HistoryOrder{
		{
			Id:                  1,
			ClientName:          client.ClientName,
			ExchangeName:        client.ExchangeName,
			Label:               "Order 1",
			Pair:                client.Pair,
			Side:                "buy",
			OrderType:           "market",
			BaseQty:             1.0,
			Price:               100.0,
			AlgorithmNamePlaced: "alg1",
			LowestSellPrc:       105.0,
			HighestBuyPrc:       95.0,
			CommissionQuoteQty:  0.1,
		},
		{
			Id:                  2,
			ClientName:          client.ClientName,
			ExchangeName:        client.ExchangeName,
			Label:               "Order 2",
			Pair:                client.Pair,
			Side:                "sell",
			OrderType:           "limit",
			BaseQty:             2.0,
			Price:               200.0,
			AlgorithmNamePlaced: "alg2",
			LowestSellPrc:       210.0,
			HighestBuyPrc:       190.0,
			CommissionQuoteQty:  0.2,
		},
	}, nil
}

func (m *mockOrderHistory) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	// Простая имитация сохранения заказа.
	if order.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}

func TestGetOrderHistory(t *testing.T) {
	mockRepo := &mockOrderHistory{}
	service := NewOrderHistoryService(mockRepo)

	client := &models.Client{
		ClientName:   "TestClient",
		ExchangeName: "TestExchange",
		Label:        "TestLabel",
		Pair:         "BTC/USD",
	}

	history, err := service.GetOrderHistory(client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(history) != 2 {
		t.Fatalf("expected 2 history orders, got %d", len(history))
	}

	asks := &models.HistoryOrder{
		Id:                  1,
		ClientName:          client.ClientName,
		ExchangeName:        client.ExchangeName,
		Label:               "Order 1",
		Pair:                client.Pair,
		Side:                "buy",
		OrderType:           "market",
		BaseQty:             1.0,
		Price:               100.0,
		AlgorithmNamePlaced: "alg1",
		LowestSellPrc:       105.0,
		HighestBuyPrc:       95.0,
		CommissionQuoteQty:  0.1,
	}
	bids := &models.HistoryOrder{
		Id:                  2,
		ClientName:          client.ClientName,
		ExchangeName:        client.ExchangeName,
		Label:               "Order 2",
		Pair:                client.Pair,
		Side:                "sell",
		OrderType:           "limit",
		BaseQty:             2.0,
		Price:               200.0,
		AlgorithmNamePlaced: "alg2",
		LowestSellPrc:       210.0,
		HighestBuyPrc:       190.0,
		CommissionQuoteQty:  0.2,
	}

	if !reflect.DeepEqual(history[0], asks) {
		t.Errorf("expected history[0] %+v, got %+v", asks, history[0])
	}
	if !reflect.DeepEqual(history[1], bids) {
		t.Errorf("expected history[1] %+v, got %+v", bids, history[1])
	}
}

func TestSaveOrder(t *testing.T) {
	mockRepo := &mockOrderHistory{}
	service := NewOrderHistoryService(mockRepo)

	client := &models.Client{
		ClientName:   "TestClient",
		ExchangeName: "TestExchange",
		Label:        "TestLabel",
		Pair:         "BTC/USD",
	}
	invalidOrder := &models.HistoryOrder{
		Price: -10.0,
	}

	err := service.SaveOrder(client, invalidOrder)
	if err == nil {
		t.Error("expected error, got nil")
	}

	validOrder := &models.HistoryOrder{
		Price: 150.0,
	}

	err = service.SaveOrder(client, validOrder)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
