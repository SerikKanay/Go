package config

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func DbConnect() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment")
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	sslmode := "disable"

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)
	fmt.Println("Connecting to DB:", dbUrl)

	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("Failed to create migrate driver:", err)
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://config/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal("Failed to create migrator:", err)
	}

	if err := migrator.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Failed to apply migrations:", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect with GORM:", err)
	}
	DB = gormDB
}
