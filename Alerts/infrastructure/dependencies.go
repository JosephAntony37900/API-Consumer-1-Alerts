package infrastructure

import (
	"log"
	"database/sql"

	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/application"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/controllers"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/repository"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/routes"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/rabbitmq"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/helpers"
	"github.com/gin-gonic/gin"
)

func InitAlertDependencies(engine *gin.Engine, db *sql.DB, rabbitmqURI string) {
	// Inicializar RabbitMQ
	if err := helpers.InitRabbitMQ(rabbitmqURI); err != nil {
		log.Fatalf("Error inicializando RabbitMQ: %v", err)
	}

	// Obtener el canal RabbitMQ
	channel := helpers.GetRabbitMQChannel()
	if channel == nil {
		log.Fatalf("RabbitMQ channel is not initialized")
	}

	// Inicializar repositorios, casos de uso y controladores
	alertRepo := repository.NewAlertRepoMySQL(db)
	createAlertUseCase := application.NewCreateAlert(alertRepo)
	getByCodigoIdAlertUseCase := application.NewGetByCodigoIdentificadorAlert(alertRepo)

	createAlertController := controllers.NewCreateAlertController(createAlertUseCase)
	getByCodigoIdAlertController := controllers.NewGetByCodigoIdentificadorAlertController(getByCodigoIdAlertUseCase)

	// Configurar rutas de Gin
	routes.SetupAlertRoutes(engine, createAlertController, getByCodigoIdAlertController)

	// Configurar el consumidor de RabbitMQ
	go func() {
		err := rabbitmq.StartAlertConsumer(createAlertUseCase, "nivel.alerta", "sensor.alerta", "amq.topic")
		if err != nil {
			log.Fatalf("Error al consumir mensajes: %v", err)
		}
	}()
}