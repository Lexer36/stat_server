package repository

import (
	"github.com/jmoiron/sqlx"
	"statServer/pkg/models"
	"statServer/pkg/repository/postgres"
)

type OrderBook interface {
	GetOrderBook(exchangeName, pair string) (*models.OrderBook, error)
	SaveOrderBook(exchangeName, pair string, asks, bids []*models.DepthOrder) error
}

type OrderHistory interface {
	GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error)
	SaveOrder(client *models.Client, order *models.HistoryOrder) error
}

type Repository struct {
	OrderBook
	OrderHistory
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		OrderBook:    postgres.NewOrderBook(db),
		OrderHistory: postgres.NewOrderHistory(db),
	}
}
