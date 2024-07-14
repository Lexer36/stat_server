package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"statServer/pkg/models"
	"testing"
)

// mock implementation of repository.OrderBook
type MockOrderBookRepo struct {
	mock.Mock
}

func (m *MockOrderBookRepo) GetOrderBook(exchangeName, pair string) (*models.OrderBook, error) {
	args := m.Called(exchangeName, pair)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.OrderBook), args.Error(1)
}

func (m *MockOrderBookRepo) SaveOrderBook(exchangeName, pair string, asks, bids []*models.DepthOrder) error {
	args := m.Called(exchangeName, pair, asks, bids)
	return args.Error(0)
}

func TestGetOrderBook(t *testing.T) {
	mockRepo := new(MockOrderBookRepo)
	service := NewOrderBookService(mockRepo)

	exchangeName := "testExchange"
	pair := "BTC/USD"
	expectedOrderBook := &models.OrderBook{
		Identity: models.Identity{ExchangeName: exchangeName, Pair: pair},
		Asks:     []*models.DepthOrder{},
		Bids:     []*models.DepthOrder{},
	}

	mockRepo.On("GetOrderBook", exchangeName, pair).Return(expectedOrderBook, nil)

	// Call the method under test
	orderBook, err := service.GetOrderBook(exchangeName, pair)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrderBook, orderBook)

	mockRepo.AssertExpectations(t)
}

func TestSaveOrderBook(t *testing.T) {
	mockRepo := new(MockOrderBookRepo)
	service := NewOrderBookService(mockRepo)

	exchangeName := "testExchange"
	pair := "BTC/USD"
	asks := []*models.DepthOrder{
		{Price: 100.0, BaseQty: 1.0},
	}
	bids := []*models.DepthOrder{
		{Price: 90.0, BaseQty: 1.5},
	}

	mockRepo.On("SaveOrderBook", exchangeName, pair, asks, bids).Return(nil)

	// Call the method under test
	err := service.SaveOrderBook(exchangeName, pair, asks, bids)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
