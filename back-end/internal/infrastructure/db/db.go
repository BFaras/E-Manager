package db

import (
    "database/sql"
    "os"
    "github.com/joho/godotenv"
    "go.uber.org/zap"
    "github.com/spf13/viper"
	"github.com/labstack/echo/v4"
)

type Server struct {
    echo *echo.Echo
    log  *zap.Logger
    cfg  *viper.Viper
    db   *sql.DB
}

func loadEnvValue(prefix string) (string, error) {
    err := godotenv.Load("../../.bin/.env")
    if err != nil {
        return "", err
    }
    return os.Getenv(prefix), nil
}

func SetUpDatabase(server *Server) (*sql.DB,error) {
    server.log.Info("Trying to get DB_URL from the environment")

    dbUrl, err := loadEnvValue("DB_URL")
    if err != nil {
        server.log.Error("Error loading DB_URL from environment: ", zap.Error(err))
        return nil,err
    }

    server.log.Info("Successfully loaded the DB URL", zap.String("DB_URL", dbUrl))

    db, err := sql.Open("postgres", dbUrl)
    if err != nil {
        server.log.Error("Failed to open database connection: ", zap.Error(err))
        return nil,err
    }

    if err := db.Ping(); err != nil {
        server.log.Error("Failed to connect to the database: ", zap.Error(err))
        db.Close()
        return nil,err
    }

    server.log.Info("Successfully connected to the database")

    var variable string
    err = db.QueryRow(`SELECT id FROM "public"."Billboard"`).Scan(&variable)
    if err != nil {
        server.log.Error("Failed to perform the query to the database: ", zap.Error(err))
        db.Close()
        return nil,err
    }

    server.log.Info("Successfully performed the query", zap.String("result", variable))
    server.log.Debug(variable)

    return db, err
}