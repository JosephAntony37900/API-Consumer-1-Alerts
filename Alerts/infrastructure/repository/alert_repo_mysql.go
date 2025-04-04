package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_"time"

	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/domain/entities"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/domain/repository"
)

type alertRepoMySQL struct {
	db *sql.DB
}

func NewAlertRepoMySQL(db *sql.DB) repository.AlertRepository {
	return &alertRepoMySQL{db: db}
}

func (repo *alertRepoMySQL) Save(alert entities.Alerts) error {
    query := "INSERT INTO Alerta (Id_Lectura, Estado, Fecha_Creacion, Codigo_Identificador, Tipo) VALUES (?, ?, ?, ?, ?)"
    _, err := repo.db.Exec(query, alert.Id_Lectura, alert.Estado, alert.Fecha_Creacion, alert.Codigo_Identificador, alert.Tipo)
    if err != nil {
        return err
    }
    return nil
}

func (repo *alertRepoMySQL) FindById(id int) (*entities.Alerts, error) {
	query := `
		SELECT Id, Id_Lectura, Estado, Fecha_Creacion, IdRol
		FROM Alerta WHERE Id = ?
	`
	row := repo.db.QueryRow(query, id)

	var alert entities.Alerts
	if err := row.Scan(&alert.Id, &alert.Id_Lectura, &alert.Estado, &alert.Fecha_Creacion, &alert.IdRol); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil 
		}
		return nil, fmt.Errorf("error obteniendo alerta por ID: %w", err)
	}

	return &alert, nil
}

func (repo *alertRepoMySQL) GetByCodigoIdentificador(codigoIdentificador string) (*entities.Alerts, error) {
	query := `
		SELECT Id, Id_Lectura, Estado, Fecha_Creacion, Codigo_Identificador, Tipo
		FROM Alerta
		WHERE Codigo_Identificador = ?
		ORDER BY Fecha_Creacion DESC
		LIMIT 1
	`

	row := repo.db.QueryRow(query, codigoIdentificador)

	var alert entities.Alerts
	if err := row.Scan(&alert.Id, &alert.Id_Lectura, &alert.Estado, &alert.Fecha_Creacion, &alert.Codigo_Identificador, &alert.Tipo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error obteniendo la alerta más reciente por Codigo_Identificador: %w", err)
	}

	return &alert, nil
}

func (repo *alertRepoMySQL) LevelReadingExists(id int) (bool, error) {
	query := `SELECT COUNT(1) FROM Lectura_Nivel WHERE Id = ?`
	var count int
	err := repo.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error verificando la existencia del nivel de lectura: %w", err)
	}
	return count > 0, nil
}

