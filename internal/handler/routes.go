package handler

import (
	"encoding/json"
	"net/http"

	"github.com/KhFirdavs/server-monitoring-go/internal/database"
	"github.com/KhFirdavs/server-monitoring-go/internal/metrics"
	"github.com/KhFirdavs/server-monitoring-go/internal/models"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/metrics", getMetrics).Methods("GET")
	return router
}
func getMetrics(w http.ResponseWriter, r *http.Request) {
	metricsData, err := metrics.CollectMetrics()
	if err != nil {
		http.Error(w, "Ошибка", 500)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metricsData)
	metricsModel := models.Metrics{
		CPUUsage:  metricsData.CPUUsage,
		RAMUsed:   metricsData.RAMUsed,
		RAMTotal:  metricsData.RAMTotal,
		DiskUsed:  metricsData.DiskUsed,
		DiskTotal: metricsData.DiskTotal,
		NetSent:   metricsData.NetSent,
		NetRecv:   metricsData.NetRecv,
	}

	db := database.NewConnectPostgres()
	err = database.SaveMetricsToDB(db, &metricsModel)
	if err != nil {
		http.Error(w, "Не удалось сохранить данные в базу данных", http.StatusInternalServerError)
		return
	}
}
