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
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)
	if err != nil {
		log.Fatal(err)
	}

	serverPort := fmt.Sprintf(":%s", cfg.Server.Port)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.SetupRouter(userHandler)

	if err := r.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}
