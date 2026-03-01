package main

import (
	"fmt"
	"gear-server/configs"
	"gear-server/internal/handler"
	repository "gear-server/internal/repository/implement"
	"gear-server/internal/router"
	service "gear-server/internal/service/implement"
	"gear-server/pkg/database"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)
	if err != nil {
		log.Fatal(err)
	}

	serverPort := fmt.Sprintf(":%s", cfg.ServerPort)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.SetupRouter(userHandler)

	if err := r.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}
