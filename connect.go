package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect to postgres db
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func connect(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}
