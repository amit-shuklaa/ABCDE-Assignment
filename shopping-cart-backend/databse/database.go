package database

import (
	"log" // log is used for log.Fatalf and log.Println

	"gorm.io/driver/postgres" // PostgreSQL driver
	"gorm.io/gorm"

	"shopping-cart-backend/models" // Adjust path based on your module name
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// PostgreSQL connection string (DSN)
	// IMPORTANT: Replace 'your_postgres_password' with the actual password you set for the 'postgres' user.
	// Ensure 'shopping_cart_db' is the name of the database you created for this project.
	// For local development, 'host' is usually 'localhost' and 'port' is '5432'.
	dsn := "host=localhost user=postgres password=Laa017@@ dbname=shopping_cart_db port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	// Open the database connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the models
	// GORM will create tables based on your structs if they don't exist,
	// or add missing columns if structs change. It won't drop columns.
	err = DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.CartItem{}, &models.Order{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
	log.Println("Database migration completed!")
}
