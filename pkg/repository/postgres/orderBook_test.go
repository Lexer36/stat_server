package postgres

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"log"
	"statServer/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrderBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error '%s' while opening mock db conn", err)
	}
	defer db.Close()

	orderBookRepo := NewOrderBook(sqlx.NewDb(db, "sqlmock"))
	exchangeName := "exchange1"
	pair := "BTC/USD"

	// Expected rows from the database
	rowsAsks := sqlmock.NewRows([]string{"price", "base_qty"}).
		AddRow(100.0, 1.0).
		AddRow(101.0, 2.0)

	rowsBids := sqlmock.NewRows([]string{"price", "base_qty"}).
		AddRow(100.0, 1.0).
		AddRow(101.0, 2.0)
	// Mocking the queries
	mock.ExpectQuery(fmt.Sprintf("SELECT price, base_qty FROM asks WHERE order_book_id =")).
		WithArgs(exchangeName, pair).
		WillReturnRows(rowsAsks)

	mock.ExpectQuery(fmt.Sprintf("SELECT price, base_qty FROM bids WHERE order_book_id =")).
		WithArgs(exchangeName, pair).
		WillReturnRows(rowsBids)

	// Call the method under test
	result, err := orderBookRepo.GetOrderBook(exchangeName, pair)
	if err != nil {
		t.Errorf("error fetching orderbook: %v", err)
		return
	}

	// Verify the result
	expected := &models.OrderBook{
		Asks: []*models.DepthOrder{
			{Price: 100.0, BaseQty: 1.0},
			{Price: 101.0, BaseQty: 2.0},
		},
		Bids: []*models.DepthOrder{
			{Price: 100.0, BaseQty: 1.0},
			{Price: 101.0, BaseQty: 2.0},
		},
		Identity: models.Identity{ExchangeName: exchangeName, Pair: pair},
	}

	assert.Equal(t, expected, result)

	// Check if all expected queries were executed
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestSaveOrderBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error '%s' while opening mock db conn", err)
	}
	defer db.Close()

	orderBookRepo := NewOrderBook(sqlx.NewDb(db, "sqlmock"))
	exchangeName := "exchange1"
	pair := "BTC/USD"
	asks := []*models.DepthOrder{
		{Price: 100.0, BaseQty: 1.0},
		{Price: 101.0, BaseQty: 2.0},
	}
	bids := []*models.DepthOrder{
		{Price: 102.0, BaseQty: 3.0},
		{Price: 103.0, BaseQty: 4.0},
	}

	// Expectations for SQL queries
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO order_book").
		WithArgs(exchangeName, pair).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT id FROM order_book").
		WithArgs(exchangeName, pair).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	for _, ask := range asks {
		mock.ExpectExec("INSERT INTO asks").
			WithArgs(1, ask.Price, ask.BaseQty).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	for _, bid := range bids {
		mock.ExpectExec("INSERT INTO bids").
			WithArgs(1, bid.Price, bid.BaseQty).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	mock.ExpectCommit()

	err = orderBookRepo.SaveOrderBook(exchangeName, pair, asks, bids)
	if err != nil {
		t.Errorf("error saving orderbook: %v", err)
		return
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
}
