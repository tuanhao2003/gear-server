package main

import (
	"fmt"
	"log"
	"net/http"

	"gear-server/configs"
	"gear-server/pkg/database"
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

	// serverPort := fmt.Sprintf(":%s", cfg.Server.Port)

	// userRepo := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

	// r := router.Setup(userHandler)

	// http.ListenAndServe(serverPort, r)
}
