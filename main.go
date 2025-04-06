package main

import (
	"log"
	"fmt"
	"os"

	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/helpers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPassword, rabbitmqHost, rabbitmqPort)

	db, err := helpers.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error inicializando la conexión a MySQL: %v", err)
	}
	defer db.Close() 

	engine := gin.Default()
	engine.Use(helpers.SetupCORS())

	infrastructure.InitAlertDependencies(engine, db, rabbitmqURI)

	engine.Run("0.0.0.0:8001")
}