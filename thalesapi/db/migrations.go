package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) *gorm.DB {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// create `products` table
			ID: "products-entity_init",
			Migrate: func(tx *gorm.DB) error {
				// it's a good practice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time
				type Product struct {
					ID           string `gorm:"primarykey"`
					CreatedAt    time.Time
					UpdatedAt    time.Time
					DeletedAt    sql.NullTime
					Name         string
					ModelNo      string
					Year         int
					ThemeType    string
					CategoryType string
					ImageURL     string
					Price        float64
					Description  sql.NullString
				}
				return tx.Migrator().CreateTable(&Product{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("products")
			},
		},
		{
			ID: "run_products_sql_script",
			Migrate: func(tx *gorm.DB) error {
				// Read the SQL file
				sqlBytes, err := os.ReadFile("./datasets/products.sql")
				if err != nil {
					return err
				}
				sqlScript := string(sqlBytes)
				// Execute the SQL script
				return tx.Exec(sqlScript).Error
			},
			Rollback: func(tx *gorm.DB) error {

				return nil
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	logrus.Println("Migration did run successfully")
	return db
}
