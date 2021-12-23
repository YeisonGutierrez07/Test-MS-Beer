package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TestMSBeerDbEngine struct {
	DB *gorm.DB
}

func NewTestMSBeerDB() (*TestMSBeerDbEngine, error) {

	// Tomamos las variables de entorno para conectar a la DB
	DbHost := os.Getenv("DB_HOST")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbUser := os.Getenv("DB_USER")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		DbHost,
		DbUser,
		DbPassword,
		DbName,
		DbPort)

	println(dsn)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return &TestMSBeerDbEngine{
		DB: db,
	}, nil

}

// Close DB
func (q *TestMSBeerDbEngine) Close() error {
	return q.Close()
}
