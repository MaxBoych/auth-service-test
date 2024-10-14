package main

import (
	"auth-service-test/internal/config"
	"auth-service-test/internal/server"
	"auth-service-test/pkg/logger"
	_ "auth-service-test/swagger/doc"
	"log"
)

func main() {
	cfg := new(config.Config)
	if err := config.LoadConfig(cfg); err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	log.Println("Config was loaded")

	if err := logger.Init(); err != nil {
		log.Fatalf("Logger init: %v", err)
	}

	s := server.NewServer(cfg)
	if err := s.Run(); err != nil {
		logger.Errorf("Cannot start server: %v", err)
		return
	}
}
