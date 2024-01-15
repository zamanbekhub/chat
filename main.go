package main

import (
	_ "chat/docs"
	"chat/internal/app"
	"chat/internal/config"
)

//go:generate go run github.com/swaggo/swag/cmd/swag init

// @title chat
// @версия 1.0.0
// @description chat
//
// @host 127.0.0.1:8001
// @BasePath /chat
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
