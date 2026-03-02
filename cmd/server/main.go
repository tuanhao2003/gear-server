package main

import (
	"fmt"
	"gear-server/config"
	"gear-server/internal/handler"
	"gear-server/internal/repository/implement"
	"gear-server/internal/router"
	"gear-server/internal/service/implement"
	"gear-server/pkg/database"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dataSource := cfg.GetDataSource()

	database.RunMigrations(dataSource)

	db, err := database.Connect(dataSource)
	if err != nil {
		log.Fatal(err)
	}

	serverPort := fmt.Sprintf(":%s", cfg.ServerPort)

	userRepo := repositoryimpl.NewUserRepository(db)
	authService := serviceimpl.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := router.SetupRouter(authHandler)

	if err := r.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}
