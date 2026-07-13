package main

import (
	"log"
	"net/http"

	"github.com/adullahkapadia/auth-service/internal/config"
	"github.com/adullahkapadia/auth-service/internal/database"
	"github.com/adullahkapadia/auth-service/internal/handler"
	"github.com/adullahkapadia/auth-service/internal/repository"
	"github.com/adullahkapadia/auth-service/internal/router"
	"github.com/adullahkapadia/auth-service/internal/service"
)


func main() {

	cfg := config.Load()

	err := database.Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate()

	if err != nil {
		log.Fatal(err)
	}

	userRepo := &repository.UserRepository{}

	authService := &service.AuthService{
		Repo: userRepo,
	}

	authHandler := &handler.AuthHandler{
		Service: authService,
		Config : cfg,
	}

	r := router.Setup(authHandler)
	
	log.Println("Server Started")

	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
	

	

}