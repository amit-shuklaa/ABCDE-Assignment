package database

import (
	"log" // Used for log.Fatalf and log.Println

	"gorm.io/driver/postgres" // PostgreSQL driver
	"gorm.io/gorm"            // GORM ORM library

	"shopping-cart-backend/models" // Path to your defined Go models
)

// DB is the global variable that holds the GORM database connection.
var DB *gorm.DB

// ConnectDatabase establishes a connection to the PostgreSQL database
// and performs automatic migrations for the defined models.
func ConnectDatabase() {
	var err error

	// PostgreSQL connection string (DSN - Data Source Name)
	// IMPORTANT:
	// - Replace 'your_postgres_password' with the actual password you set for the 'postgres' user.
	// - Ensure 'shopping_cart_db' is the exact name of the database you created for this project in PostgreSQL.
	// - For local development, 'host' is usually 'localhost' and 'port' is '5432'.
	// - TimeZone is set to Asia/Kolkata as per your previous location context.
	dsn := "host=localhost user=postgres password=Laa017@@ dbname=shopping_cart_db port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	// Open the database connection using the PostgreSQL driver
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// If connection fails, log a fatal error and exit the application
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the models
	// GORM will automatically create tables in your database based on your
	// struct definitions (models.User, models.Item, etc.) if they don't exist.
	// It will also add missing columns if you modify your structs, but won't drop existing columns.
	err = DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.CartItem{}, &models.Order{})
	if err != nil {
		// If migration fails, log a fatal error
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
	log.Println("Database migration completed!")
}
