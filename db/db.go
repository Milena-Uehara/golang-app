package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func ConnectDb() *sql.DB {
	connStr := os.ExpandEnv("user=${DB_USER} dbname=${DB_NAME} host=${DB_HOST} password=${DB_PASSWORD} sslmode=disable")
	//connStr := "user=postgres dbname=products_db host=localhost password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		zap.L().Error(err.Error())
	}
	return db
}
