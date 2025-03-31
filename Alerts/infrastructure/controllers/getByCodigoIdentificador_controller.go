package controllers

import (
	"log"
	"net/http"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/application"
	"github.com/gin-gonic/gin"
)

type GetByCodigoIdentificadorAlertController struct {
	getUseCase *application.GetByCodigoIdentificadorAlert
}

func NewGetByCodigoIdentificadorAlertController(getUseCase *application.GetByCodigoIdentificadorAlert) *GetByCodigoIdentificadorAlertController {
	return &GetByCodigoIdentificadorAlertController{getUseCase: getUseCase}
}

func (c *GetByCodigoIdentificadorAlertController) Handle(ctx *gin.Context) {
	codigoIdentificador := ctx.Query("codigo_identificador")
	if codigoIdentificador == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'codigo_identificador' es requerido"})
		return
	}

	alert, err := c.getUseCase.Run(codigoIdentificador)
	if err != nil {
		log.Printf("Error obteniendo la última alerta: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la última alerta"})
		return
	}

	if alert == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay alertas disponibles para este identificador"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"alert": alert})
}