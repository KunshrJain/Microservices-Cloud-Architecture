package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Order struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Item   string  `json:"item"`
	Amount float64 `json:"amount"`
}

var orders = []Order{}

func main() {
	http.HandleFunc("/orders", handleOrders)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println("Order Service running on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func handleOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var o Order
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		o.ID = len(orders) + 1
		orders = append(orders, o)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(o)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}
