package database

import (
	"fmt"

	"github.com/KhFirdavs/server-monitoring-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectPostgres() *gorm.DB {
	Host := "localhost"
	Port := "5432"
	User := "postgres"
	Password := "77887788"
	dbName := "postgres"
	dbParams := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, dbName)
	db, err := gorm.Open(postgres.Open(dbParams))
	if err != nil {
		panic("Ошибка подключения к базе данных")
	}
	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Metrics{})
	if err != nil {
		panic("Ошибка автосохранения")
	}
}

func SaveMetricsToDB(db *gorm.DB, metrics *models.Metrics) error {
	if err := db.Create(metrics).Error; err != nil {
		return fmt.Errorf("failed to save metrics to db: %w", err)
	}
	return nil
}
