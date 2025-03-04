package dbase

import (
	"Employee/internal/config"
	"Employee/internal/employee/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPostgersDB(cfg config.AppConfig) (*gorm.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name,
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err = db.AutoMigrate(&entity.Employee{}); err != nil {
		log.Fatalf("error auto migrate employee: %v", err)
	}

	return db, err
}
