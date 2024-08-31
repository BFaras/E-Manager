package db

import (
	"back-end/internal/infrastructure/logger"
	"database/sql"
	"os"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
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

    var variable string
    err = db.QueryRow(`SELECT id FROM "public"."Billboard"`).Scan(&variable)
    if err != nil {
        logger.Error("Failed to perform the query to the database: ", zap.Error(err))
        db.Close()
        return nil, err
    }

    logger.Info("Successfully performed the query", zap.String("result", variable))
    logger.Debug(variable)

    return db, nil
}
