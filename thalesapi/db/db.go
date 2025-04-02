package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
	"thalesapi/configs"
)

var (
	dbMain *gorm.DB
	dbOnce sync.Once
)

// InitDB Provides *gorm.DB for general use and migration
func InitDB(c *configs.DBConfig) *gorm.DB {

	dbOnce.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Password, c.Name)

		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect db: %v", err)
		}

		dbMain = RunMigrations(gormDB)

	})

	return dbMain
}
