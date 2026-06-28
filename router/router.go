package router

import (
	"go-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("PUT", "OPTIONS")

	return r
}
