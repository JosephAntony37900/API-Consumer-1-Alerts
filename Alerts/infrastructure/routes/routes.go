package routes

import (
	controllers "github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/controllers"
	"github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/websocket"
)

func SetupAlertRoutes(engine *gin.Engine, createAlertController *controllers.CreateAlertController, getByCodigoIdentificadorAlertController *controllers.GetByCodigoIdentificadorAlertController) {
	engine.POST("/alerts", createAlertController.Handle)
	engine.GET("/alerts/latest", getByCodigoIdentificadorAlertController.Handle)
	engine.GET("/ws", websocket.HandleConnections)
}