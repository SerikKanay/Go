package config

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func DbConnect() {
	dsn := os.Getenv("DB")
	sqlDb, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := migratepg.WithInstance(sqlDb, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://config/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := migrator.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = gormDB
}
