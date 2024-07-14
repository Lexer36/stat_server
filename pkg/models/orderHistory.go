package models

import "time"

type HistoryOrder struct {
	Id                  int       `db:"id"`
	ClientName          string    `json:"client_name" db:"client_name"`
	ExchangeName        string    `json:"exchange_name" db:"exchange_name"`
	Label               string    `json:"label" db:"label"`
	Pair                string    `json:"pair" db:"pair"`
	Side                string    `json:"side" db:"side"`
	OrderType           string    `json:"order_type" db:"type"`
	BaseQty             float64   `json:"base_qty" db:"base_qty"`
	Price               float64   `json:"price" db:"price"`
	AlgorithmNamePlaced string    `json:"algorithm_name_placed" db:"algorithm_name_placed"`
	LowestSellPrc       float64   `json:"lowest_sell_prc" db:"lowest_sell_prc"`
	HighestBuyPrc       float64   `json:"highest_buy_prc" db:"highest_buy_prc"`
	CommissionQuoteQty  float64   `json:"commission_quote_qty" db:"commission_quote_qty"`
	TimePlaced          time.Time `json:"time_placed" db:"time_placed"`
}

type Client struct {
	ClientName   string `json:"client_name"`
	ExchangeName string `json:"exchange_name"`
	Label        string `json:"label"`
	Pair         string `json:"pair"`
}
