package main

import (
	"net/http"

	"github.com/Milena-Uehara/golang-app/routes"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	routes.LoadRoutes()

	port := ":8000"

	zap.S().Infof("Server started on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		zap.L().Fatal("Failed to start server", zap.Error(err))
	}
}
