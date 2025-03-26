package routes

import (
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAlertRoutes(engine *gin.Engine, createAlertController *controllers.CreateAlertController, getByUserIdAlertController *controllers.GetByUserIdAlertController) {
	// Ruta para crear alertas
	engine.POST("/alerts", createAlertController.Handle)

	// Ruta para obtener alertas por ID de rol de usuario
	engine.GET("/alerts/:id", getByUserIdAlertController.Handle)
}