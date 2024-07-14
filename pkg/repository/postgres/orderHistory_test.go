package postgres

import (
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"statServer/pkg/models"
)

func TestGetOrderHistory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	orderHistoryRepo := NewOrderHistory(sqlx.NewDb(db, "sqlmock"))
	client := &models.Client{
		ClientName:   "testClient",
		ExchangeName: "testExchange",
		Label:        "testLabel",
		Pair:         "BTC/USD",
	}
	timeT := time.Now()
	// Expected rows from the database
	rows := sqlmock.NewRows([]string{
		"id", "client_name", "exchange_name", "label", "pair", "side", "type", "base_qty", "price",
		"algorithm_name_placed", "lowest_sell_prc", "highest_buy_prc", "commission_quote_qty", "time_placed",
	}).AddRow(1, "testClient", "testExchange", "testLabel", "BTC/USD", "buy", "limit", 1.0, 100.0,
		"algo1", 90.0, 110.0, 0.1, timeT)

	// Expect the query with arguments and return rows
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM order_history WHERE client_name = $1 AND exchange_name = $2 AND label = $3 AND pair = $4")).
		WithArgs(client.ClientName, client.ExchangeName, client.Label, client.Pair).
		WillReturnRows(rows)

	// Call the method under test
	result, err := orderHistoryRepo.GetOrderHistory(client)
	if err != nil {
		t.Errorf("error fetching order history: %v", err)
		return
	}

	// Verify the result
	expected := []*models.HistoryOrder{
		{
			Id:                  1,
			ClientName:          "testClient",
			ExchangeName:        "testExchange",
			Label:               "testLabel",
			Pair:                "BTC/USD",
			Side:                "buy",
			OrderType:           "limit",
			BaseQty:             1.0,
			Price:               100.0,
			AlgorithmNamePlaced: "algo1",
			LowestSellPrc:       90.0,
			HighestBuyPrc:       110.0,
			CommissionQuoteQty:  0.1,
			TimePlaced:          timeT,
		},
	}

	assert.Equal(t, expected, result)

	// Check if all expected queries were executed
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	orderHistoryRepo := NewOrderHistory(sqlx.NewDb(db, "sqlmock"))
	client := &models.Client{
		ClientName:   "testClient",
		ExchangeName: "testExchange",
		Label:        "testLabel",
		Pair:         "BTC/USD",
	}
	order := &models.HistoryOrder{
		Side:                "buy",
		OrderType:           "limit",
		BaseQty:             1.0,
		Price:               100.0,
		AlgorithmNamePlaced: "algo1",
		LowestSellPrc:       90.0,
		HighestBuyPrc:       110.0,
		CommissionQuoteQty:  0.1,
		TimePlaced:          time.Now(),
	}

	// Expectations for SQL query
	mock.ExpectExec("INSERT INTO order_history").
		WithArgs(client.ClientName, client.ExchangeName, client.Label, client.Pair,
			order.Side, order.OrderType, order.BaseQty, order.Price, order.AlgorithmNamePlaced,
								order.LowestSellPrc, order.HighestBuyPrc, order.CommissionQuoteQty, order.TimePlaced).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Mocking insert success

	// Call the method under test
	err = orderHistoryRepo.SaveOrder(client, order)
	if err != nil {
		t.Errorf("error saving order: %v", err)
		return
	}

	// Check if all expected queries were executed
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
