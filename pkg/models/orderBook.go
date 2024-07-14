package models

type DepthOrder struct {
	Price   float64 `db:"price"`
	BaseQty float64 `db:"base_qty"`
}

type Identity struct {
	ExchangeName string `json:"exchangeName"`
	Pair         string `json:"pair"`
}

type OrderBook struct {
	Identity
	Asks []*DepthOrder `json:"asks"`
	Bids []*DepthOrder `json:"bids"`
}
