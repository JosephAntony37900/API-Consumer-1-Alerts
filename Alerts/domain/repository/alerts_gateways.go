package repository

import "github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/domain/entities"

type AlertRepository interface {
	Save(alert entities.Alerts) error
	FindById(id int) (*entities.Alerts, error)
	GetByUserId(id int) (*entities.Alerts, error)
	LevelReadingExists(id int) (bool, error)
}