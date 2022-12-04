package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type DatabaseConfig struct {
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_DRIVER   string
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_DRIVER:   os.Getenv("POSTGRES_DRIVER"),
	}
}

func Init() *gorm.DB {
	dbConfig := getDatabaseConfig()
	// Originally should be set as "postgres". In GCP env
	// the driver name is "cloudsqlpostgres"
	driverName := dbConfig.POSTGRES_DRIVER
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.POSTGRES_HOST,
		dbConfig.POSTGRES_USER,
		dbConfig.POSTGRES_PASSWORD,
		dbConfig.POSTGRES_DB,
		dbConfig.POSTGRES_PORT,
	)

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DriverName: driverName,
			DSN:        dsn,
		}),
		&gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Silent),
		})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", dbConfig.POSTGRES_DB))

	return db
}
