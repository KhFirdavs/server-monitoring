package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/metrics", getMetrics).Methods("GET")
	return router
}
func getMetrics(w http.ResponseWriter, r *http.Request) {
	// Your logic to collect and return metrics goes here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Metrics data"))
}
