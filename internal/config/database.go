package config

import (
	"fmt"
	"log"

	"kikukafandi/book-shop-api/internal/adapter/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig holds database configuration.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewDatabase creates a new database connection.
func NewDatabase(cfg DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")
	return database, nil
}

// AutoMigrate runs auto migration for all models.
func AutoMigrate(database *gorm.DB) error {
	return database.AutoMigrate(
		&db.BookModel{},
		&db.UserModel{},
		&db.OrderModel{},
	)
}
