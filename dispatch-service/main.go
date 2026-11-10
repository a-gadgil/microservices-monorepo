package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RideRequest struct {
	RiderID     string  `json:"rider_id"`
	PickupLat   float64 `json:"pickup_lat"`
	PickupLong  float64 `json:"pickup_long"`
	Destination string  `json:"destination"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "UP",
		"service": "dispatch-service",
	})
}

func dispatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RideRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Ride requested successfully. Searching for nearby drivers...",
		"ride_id":  "ride_abc123",
		"rider_id": req.RiderID,
		"status":   "DISPATCHING",
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/rides", dispatchHandler)

	port := ":8080"
	fmt.Printf("Dispatch Service running on port %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
