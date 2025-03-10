package db

import (
	"back-end/internal/infrastructure/logger"
	"database/sql"
	"os"
	"github.com/joho/godotenv"
)


func loadEnvValue(prefix string) (string, error) {
    err := godotenv.Load("../../.bin/.env")
    if err != nil {
        return "", err
    }
    return os.Getenv(prefix), nil
}

func SetUpDatabase() (*sql.DB, error) {
    dbUrl, err := loadEnvValue("DB_URL")
    if err != nil {
        return nil, err
    }

    db, err := sql.Open("postgres", dbUrl)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    logger.Info("Successfully connected to the database")

    return db, nil
}
