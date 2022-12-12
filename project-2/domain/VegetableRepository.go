package domain

type VegetableRepository interface {
	GetAll() ([]Vegetable, error)
}