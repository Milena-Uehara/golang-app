package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func ConnectDb() *sql.DB {
	connStr := "user=postgres dbname=alura_loja host=localhost password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		zap.L().Error(err.Error())
	}
	return db
}
