package main

import (
	"log"

	"github.com/ncondezo/final/cmd/server/router"
	"github.com/ncondezo/final/docs"
	"github.com/ncondezo/final/pkg/store"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

const (
	serverPort = ":8080"
)

// @title Desafío II - Backend Go
// @version 1.0
// @description API para la gestión de turnos de una clínica dental.
func main() {

	store.NewMySQLConnection()
	database := store.GetConnection()

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	routes := router.NewRouter(engine, database)
	routes.BuildRoutes()

	docs.SwaggerInfo.BasePath = "localhost:8080/api/v1"

	if engineError := engine.Run(serverPort); engineError != nil {
		log.Fatal(engineError)
	}

}
