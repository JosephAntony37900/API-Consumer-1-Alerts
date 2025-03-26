package application

import (
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/domain/entities"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/domain/repository"
)

type GetByUserIdAlert struct {
	repo repository.AlertRepository
}

func NewGetByUserIdAlert(repo repository.AlertRepository) *GetByUserIdAlert{
	return &GetByUserIdAlert{repo: repo}
}

func (gbuia *GetByUserIdAlert) Run(Id int) ( *entities.Alerts ,error) {
	alert, err := gbuia.repo.GetByUserId(Id)
	if err != nil {
		return nil, err
	}

	return alert,nil
}