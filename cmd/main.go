package main

import (
	"log"

	"github.com/KhFirdavs/server-monitoring-go/internal/api"
	"github.com/KhFirdavs/server-monitoring-go/internal/database"
	"github.com/KhFirdavs/server-monitoring-go/internal/handler"
	"github.com/KhFirdavs/server-monitoring-go/internal/metrics"
	"github.com/KhFirdavs/server-monitoring-go/internal/models"
)

func main() {
	metrics.StartCollector()
	log.Println("Сбор метрик запущен...")

	router := handler.NewRouter()

	srv := &api.Server{}
	if err := srv.ServerRun(router, "8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %s", err.Error())
	}

	db := database.NewConnectPostgres()
	if err := db.AutoMigrate(&models.Metrics{}); err != nil {
		log.Fatalf("Ошибка переноса базы данных: %s", err.Error())
	}
}
