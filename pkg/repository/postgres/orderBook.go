package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"statServer/pkg/models"
)

type OrderBook struct {
	db *sqlx.DB
}

func NewOrderBook(db *sqlx.DB) *OrderBook {
	return &OrderBook{db: db}
}

func (r *OrderBook) GetOrderBook(exchangeName, pair string) (*models.OrderBook, error) {
	var asks []*models.DepthOrder
	var bids []*models.DepthOrder
	query := `SELECT price, base_qty FROM asks 
         WHERE order_book_id = (SELECT id FROM order_book 
								  WHERE exchange = $1 AND pair = $2)`
	err := r.db.Select(&asks, query, exchangeName, pair)
	if err != nil {
		return nil, err
	}
	query = `SELECT price, base_qty FROM bids 
         WHERE order_book_id = (SELECT id FROM order_book 
								  WHERE exchange = $1 AND pair = $2)`
	err = r.db.Select(&bids, query, exchangeName, pair)
	if err != nil {
		return nil, err
	}

	return &models.OrderBook{
		Bids:     bids,
		Asks:     asks,
		Identity: models.Identity{ExchangeName: exchangeName, Pair: pair},
	}, nil
}

func (r *OrderBook) SaveOrderBook(exchangeName, pair string, asks, bids []*models.DepthOrder) error {
	// Create a transaction
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %w", err)
	}

	defer func() {
		// Rollback the transaction if there's an error
		if err != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to error: %v", err)
		} else {
			// Commit the transaction if successful
			err = tx.Commit()
			if err != nil {
				log.Printf("Error committing transaction: %v", err)
			}
		}
	}()

	// Insert into order_book table
	_, err = tx.Exec(`
        INSERT INTO order_book (exchange, pair)
        VALUES ($1, $2)
        ON CONFLICT (exchange, pair) DO NOTHING`,
		exchangeName, pair)
	if err != nil {
		return fmt.Errorf("error inserting into order_book table: %w", err)
	}

	// Fetch order_book_id for the inserted or existing record
	var orderBookID int
	err = tx.QueryRow(`
        SELECT id FROM order_book
        WHERE exchange = $1 AND pair = $2
        LIMIT 1`,
		exchangeName, pair).Scan(&orderBookID)
	if err != nil {
		return fmt.Errorf("error fetching order_book_id: %w", err)
	}

	// Insert asks into asks table
	for _, ask := range asks {
		_, err = tx.Exec(`
            INSERT INTO asks (order_book_id, price, base_qty)
            VALUES ($1, $2, $3)`,
			orderBookID, ask.Price, ask.BaseQty)
		if err != nil {
			return fmt.Errorf("error inserting ask: %w", err)
		}
	}

	// Insert bids into bids table
	for _, bid := range bids {
		_, err = tx.Exec(`
            INSERT INTO bids (order_book_id, price, base_qty)
            VALUES ($1, $2, $3)`,
			orderBookID, bid.Price, bid.BaseQty)
		if err != nil {
			return fmt.Errorf("error inserting bid: %w", err)
		}
	}

	return nil
}
