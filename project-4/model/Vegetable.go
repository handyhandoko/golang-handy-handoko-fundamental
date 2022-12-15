package model

type Vegetable struct {
	Id    uint
	Name  string `validate:"required,min=4"`
	Price uint   `validate:"required,min=1000"`
}
