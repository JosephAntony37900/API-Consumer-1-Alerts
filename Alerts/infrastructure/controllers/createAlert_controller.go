package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/application"
	"github.com/gin-gonic/gin"
)

type CreateAlertController struct {
	createUseCase *application.CreateAlert
}

func NewCreateAlertController(createUseCase *application.CreateAlert) *CreateAlertController {
	return &CreateAlertController{createUseCase: createUseCase}
}

func (c *CreateAlertController) Handle(ctx *gin.Context) {
	var request struct {
		Id_Lectura int    `json:"id_lectura"`
		Estado     string `json:"estado"`
		IdRol      int    `json:"id_rol"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error en el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	

	fechaCreacion := time.Now()
	if err := c.createUseCase.Run(request.Id_Lectura, request.Estado, fechaCreacion); err != nil {
		log.Printf("Error creando la alerta: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la alerta"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Alerta creada exitosamente"})
}