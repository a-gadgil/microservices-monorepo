package main

import (
	"encoding/json"
	"net/http"
)

type InventoryItem struct {
	ItemID   string `json:"item_id"`
	StockQty int    `json:"stock_qty"`
	Status   string `json:"status"`
}

func main() {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "UP",
			"service": "inventory-service",
		})
	})

	// Inventory stock check endpoint
	mux.HandleFunc("/api/inventory", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(InventoryItem{
			ItemID:   "ITEM-MILK-01",
			StockQty: 45,
			Status:   "IN_STOCK",
		})
	})

	port := ":8002"
	println("Inventory Service running on port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}