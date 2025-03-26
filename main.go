package main

import (
	"log"
	"fmt"
	"os"
	_"time"

	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/application"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/controllers"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/repository"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/routes"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/rabbitmq"
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

	if err := helpers.InitRabbitMQ(rabbitmqURI); err != nil {
		log.Fatalf("Error inicializando RabbitMQ: %v", err)
	}
	defer helpers.CloseRabbitMQ()

	// Repositorios
	alertRepo := repository.NewAlertRepoMySQL(db)

	// Casos de uso
	createAlertUseCase := application.NewCreateAlert(alertRepo)
	getByUserIdAlertUseCase := application.NewGetByUserIdAlert(alertRepo)

	// Controladores
	createAlertController := controllers.NewCreateAlertController(createAlertUseCase)
	getByUserIdAlertController := controllers.NewGetByUserIdAlertController(getByUserIdAlertUseCase)

	engine := gin.Default()
	routes.SetupAlertRoutes(engine, createAlertController, getByUserIdAlertController)

	// Iniciar consumo de mensajes desde RabbitMQ
	go func() {
		err = rabbitmq.StartAlertConsumer(createAlertUseCase, "nivel.alerta", "sensor.alerta", "amq.topic")
		if err != nil {
			log.Fatalf("Error al consumir mensajes: %v", err)
		}
	}()

	engine.Use(helpers.SetupCORS())
	engine.Run(":8001")
}