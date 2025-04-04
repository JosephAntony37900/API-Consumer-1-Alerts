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
	// Cargar las variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	// Configuración de RabbitMQ
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPassword, rabbitmqHost, rabbitmqPort)

	// Inicializar conexión a MySQL
	db, err := helpers.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error inicializando la conexión a MySQL: %v", err)
	}
	defer db.Close() // Cerrar la conexión al final de la ejecución

	// Inicializar Gin y configuración
	engine := gin.Default()
	engine.Use(helpers.SetupCORS())

	// Inicializar las dependencias de alertas
	infrastructure.InitAlertDependencies(engine, db, rabbitmqURI)

	// Iniciar el servidor
	engine.Run(":8001")
}