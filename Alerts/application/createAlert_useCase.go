package application

import (
	"time"
	"fmt"
	"log"

	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/domain/entities"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/domain/repository"
)

type CreateAlert struct {
	repo repository.AlertRepository
}

func NewCreateAlert(repo repository.AlertRepository) *CreateAlert{
	return &CreateAlert{repo: repo}
}

func (cal *CreateAlert) Run(Id_Lectura int, Estado string, Fecha_Creacion time.Time, Codigo_Identificador string, Tipo bool) error {
    // Verificar si el Id_Lectura existe en la tabla Lectura_Nivel
    exists, err := cal.repo.LevelReadingExists(Id_Lectura)
    if err != nil {
        return fmt.Errorf("error verificando la existencia del nivel de lectura: %w", err)
    }
    if !exists {
        log.Printf("ID de lectura inválido en el mensaje de alerta: %d", Id_Lectura)
        return fmt.Errorf("ID de lectura inválido")
    }

    alert := entities.Alerts{
        Id_Lectura:          Id_Lectura,
        Estado:              Estado,
        Fecha_Creacion:      Fecha_Creacion,
        Codigo_Identificador: Codigo_Identificador,
        Tipo:                Tipo, // Nuevo atributo
    }

    if err := cal.repo.Save(alert); err != nil {
        return fmt.Errorf("error guardando la alerta: %w", err)
    }

    return nil
}