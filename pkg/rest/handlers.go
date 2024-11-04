package rest

import (
	"context"
	"encoding/json"
	"main.go/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
)

var ctx = context.Background()

// AddItemToCart handles adding an item to a user's cart.
func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	var itemInfo struct {
		ItemID   int `json:"item_id"`
		Quantity int `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&itemInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := storage.AddItemToCart(userID, itemInfo.ItemID, itemInfo.Quantity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Checkout calculates the total value of a user's cart.
func Checkout(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]

	cartItems, err := storage.GetCartItems(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	total, err := calculateTotal(cartItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"total": total})
}

// Pay processes payment and clears the user's cart.
func Pay(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]

	cartItems, err := storage.GetCartItems(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var paymentInfo struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&paymentInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	total, err := calculateTotal(cartItems)
	if err != nil || paymentInfo.Amount < total {
		http.Error(w, "invalid payment amount", http.StatusBadRequest)
		return
	}

	// Simulate successful payment
	err = storage.DecreaseStock(cartItems) // Decrease stock after successful payment
	if err != nil {
		http.Error(w, "failed to decrease stock", http.StatusInternalServerError)
		return
	}

	// Simulate successful payment
	if err := storage.ClearCart(userID); err != nil {
		http.Error(w, "failed to clear cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment successful!"))
}

func calculateTotal(cartItems map[int]int) (float64, error) {
	var total float64
	for itemID, qty := range cartItems {
		item, err := storage.GetItem(itemID)
		if err != nil {
			return 0, err
		}
		total += float64(qty) * item.Price
	}
	return total, nil
}
