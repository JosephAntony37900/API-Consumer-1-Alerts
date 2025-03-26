package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/application"
	"github.com/gin-gonic/gin"
)

type GetByUserIdAlertController struct {
	getUseCase *application.GetByUserIdAlert
}

func NewGetByUserIdAlertController(getUseCase *application.GetByUserIdAlert) *GetByUserIdAlertController {
	return &GetByUserIdAlertController{getUseCase: getUseCase}
}

func (c *GetByUserIdAlertController) Handle(ctx *gin.Context) {
	userIdStr := ctx.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("Error al convertir el ID: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	alert, err := c.getUseCase.Run(userId)
	if err != nil {
		log.Printf("Error obteniendo alerta: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la alerta"})
		return
	}

	if alert == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No se encontraron alertas"})
		return
	}

	ctx.JSON(http.StatusOK, alert)
}