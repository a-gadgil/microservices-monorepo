package main

import (
	"encoding/json"
	"net/http"
)

type Dispatch struct {
	OrderID   string `json:"order_id"`
	RiderName string `json:"rider_name"`
	OTP       string `json:"otp"`
	Status    string `json:"status"`
}

func main() {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "UP",
			"service": "dispatch-service",
		})
	})

	// Dispatch & OTP assignment endpoint
	mux.HandleFunc("/api/dispatch", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Dispatch{
			OrderID:   "ORD-998877",
			RiderName: "Rahul Kumar",
			OTP:       "4829",
			Status:    "ASSIGNED",
		})
	})

	port := ":8003"
	println("Dispatch Service running on port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}