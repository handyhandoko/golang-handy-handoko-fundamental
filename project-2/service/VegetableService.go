package service

import "project-2/domain"

type VegetableService interface{
	GetAllVegetables() ([]domain.Vegetable, error)
}