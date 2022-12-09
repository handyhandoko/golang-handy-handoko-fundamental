package service

import (
	"project-2/domain"
)

type DefaultVegetableService struct {
	repo domain.VegetableRepository
}

func (service DefaultVegetableService) GetAllVegetables() ([]domain.Vegetable, error){
	return service.repo.GetAll()
}

func NewVegetableService(repo domain.VegetableRepository) DefaultVegetableService{
	return DefaultVegetableService{repo}
}