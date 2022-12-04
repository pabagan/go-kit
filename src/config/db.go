package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func Init() *gorm.DB {
	// Environment variables
	ENV := os.Getenv("ENV")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	driverName := "postgres"

	// use this one in case using GCP Proxy
	if ENV != "local" {
		driverName = "cloudsqlpostgres"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT)

	db, err := gorm.Open(postgres.New(postgres.Config{DriverName: driverName, DSN: dsn}), &gorm.Config{SkipDefaultTransaction: true, Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", SCHEMA_NAME))

	return db
}
