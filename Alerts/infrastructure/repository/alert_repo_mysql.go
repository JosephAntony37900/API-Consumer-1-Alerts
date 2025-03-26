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
	query := `
		INSERT INTO Alerta (Id_Lectura, Estado, Fecha_Creacion, IdRol)
		VALUES (?, ?, ?, ?)
	`
	_, err := repo.db.Exec(query, alert.Id_Lectura, alert.Estado, alert.Fecha_Creacion, alert.IdRol)
	if err != nil {
		return fmt.Errorf("error guardando la alerta: %w", err)
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
			return nil, nil // No se encontraron alertas
		}
		return nil, fmt.Errorf("error obteniendo alerta por ID: %w", err)
	}

	return &alert, nil
}

func (repo *alertRepoMySQL) GetByUserId(id int) (*entities.Alerts, error) {
	query := `
		SELECT Id, Id_Lectura, Estado, Fecha_Creacion, IdRol
		FROM Alerta WHERE IdRol = ?
	`
	row := repo.db.QueryRow(query, id)

	var alert entities.Alerts
	if err := row.Scan(&alert.Id, &alert.Id_Lectura, &alert.Estado, &alert.Fecha_Creacion, &alert.IdRol); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No se encontraron alertas
		}
		return nil, fmt.Errorf("error obteniendo alerta por ID de usuario: %w", err)
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