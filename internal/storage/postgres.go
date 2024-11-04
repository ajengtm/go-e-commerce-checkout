package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/pkg/item"
)

var DB *pgxpool.Pool

// InitDB initializes the database connection pool.
func InitDB(connString string) error {
	var err error
	DB, err = pgxpool.New(context.Background(), connString)
	return err
}

// GetItem retrieves an item from the database.
func GetItem(id int) (item.Item, error) {
	var item item.Item
	err := DB.QueryRow(context.Background(), "SELECT id, name, price FROM items WHERE id=$1", id).
		Scan(&item.ID, &item.Name, &item.Price)
	return item, err
}

// AddItemToCart adds an item to a user's cart in the database.
func AddItemToCart(userID string, itemID, quantity int) error {
	_, err := DB.Exec(context.Background(),
		`INSERT INTO carts (user_id, item_id, quantity) 
		 VALUES ($1, $2, $3) 
		 ON CONFLICT (user_id, item_id) 
		 DO UPDATE SET quantity = carts.quantity + $3`,
		userID, itemID, quantity)
	return err
}

// DecreaseStock decreases the stock of an item after a successful payment.
func DecreaseStock(cartItems map[int]int) error {
	tx, err := DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	for itemID, quantity := range cartItems {
		// Lock the row for update and check stock
		var stock int
		err := tx.QueryRow(context.Background(),
			`SELECT stock FROM items WHERE id = $1 FOR UPDATE`,
			itemID).Scan(&stock)
		if err != nil {
			return err
		}

		if stock < quantity {
			return fmt.Errorf("not enough stock for item %d", itemID)
		}

		// Decrease the stock for each item in the cart
		_, err = tx.Exec(context.Background(),
			`UPDATE items SET stock = stock - $1 WHERE id = $2`,
			quantity, itemID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// GetCartItems retrieves all items in a user's cart.
func GetCartItems(userID string) (map[int]int, error) {
	rows, err := DB.Query(context.Background(), "SELECT item_id, quantity FROM carts WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cartItems := make(map[int]int)
	for rows.Next() {
		var itemID, quantity int
		if err := rows.Scan(&itemID, &quantity); err != nil {
			return nil, err
		}
		cartItems[itemID] = quantity
	}
	return cartItems, rows.Err()
}

// ClearCart removes all items from a user's cart.
func ClearCart(userID string) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM carts WHERE user_id=$1", userID)
	return err
}
