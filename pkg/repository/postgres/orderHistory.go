package postgres

import (
	"github.com/jmoiron/sqlx"
	"statServer/pkg/models"
)

type OrderHistory struct {
	db *sqlx.DB
}

func NewOrderHistory(db *sqlx.DB) *OrderHistory {
	return &OrderHistory{db: db}
}

func (r *OrderHistory) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	var res []*models.HistoryOrder
	query := `SELECT * FROM order_history
				WHERE client_name = $1
				AND exchange_name = $2
				AND label = $3
				AND pair = $4`
	err := r.db.Select(&res, query, client.ClientName, client.ExchangeName, client.Label, client.Pair)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *OrderHistory) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	query := `
		INSERT INTO order_history (client_name, exchange_name, label, pair, side, type, base_qty, price, algorithm_name_placed, lowest_sell_prc, highest_buy_prc, commission_quote_qty, time_placed)
		VALUES (:client_name, :exchange_name, :label, :pair, :side, :type, :base_qty, :price, :algorithm_name_placed, :lowest_sell_prc, :highest_buy_prc, :commission_quote_qty, :time_placed)
	`
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"client_name":           client.ClientName,
		"exchange_name":         client.ExchangeName,
		"label":                 client.Label,
		"pair":                  client.Pair,
		"side":                  order.Side,
		"type":                  order.OrderType,
		"base_qty":              order.BaseQty,
		"price":                 order.Price,
		"algorithm_name_placed": order.AlgorithmNamePlaced,
		"lowest_sell_prc":       order.LowestSellPrc,
		"highest_buy_prc":       order.HighestBuyPrc,
		"commission_quote_qty":  order.CommissionQuoteQty,
		"time_placed":           order.TimePlaced,
	})
	if err != nil {
		return err
	}

	return nil
}
