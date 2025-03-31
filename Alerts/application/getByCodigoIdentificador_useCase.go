package application

import (
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/domain/entities"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/domain/repository"
)

type GetByCodigoIdentificadorAlert struct {
	repo repository.AlertRepository
}

func NewGetByCodigoIdentificadorAlert(repo repository.AlertRepository) *GetByCodigoIdentificadorAlert {
	return &GetByCodigoIdentificadorAlert{repo: repo}
}

func (gbcia *GetByCodigoIdentificadorAlert) Run(codigoIdentificador string) (*entities.Alerts, error) {
	alert, err := gbcia.repo.GetByCodigoIdentificador(codigoIdentificador)
	if err != nil {
		return nil, err
	}

	return alert, nil
}