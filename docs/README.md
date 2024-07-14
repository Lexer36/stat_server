Endpoints:
1. Get order-book
 >/order-book/get

- принимаемые параметры body:
 > type Identity struct {
 ExchangeName string `json:"exchangeName"`
 Pair         string `json:"pair"`
 }

1. Save order-book
>/order-book/save

- принимаемые параметры body:
 > type DepthOrder struct {
  Price   float64 `db:"price"`
  BaseQty float64 `db:"base_qty"`
  }

 > type Identity struct {
ExchangeName string `json:"exchangeName"`
Pair         string `json:"pair"`
}

> type OrderBook struct {
Identity
Asks []*DepthOrder `json:"asks"`
Bids []*DepthOrder `json:"bids"`
}

 3. Get order-history
     >/order-history/get

- принимаемые параметры body:
> type Client struct {
ClientName   string `json:"client_name"`
ExchangeName string `json:"exchange_name"`
Label        string `json:"label"`
Pair         string `json:"pair"`
}
4. Save order-history
     >/order-history/save

- принимаемые параметры body:
> type HistoryOrder struct {
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