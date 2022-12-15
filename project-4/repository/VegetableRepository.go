package repository

import (
	"fmt"
	"project-2/model"
)

func AddItem(slice []model.Vegetable, vegetable model.Vegetable) []model.Vegetable {
	var lastItem model.Vegetable = slice[len(slice)-1]
	vegetable.Id = lastItem.Id + 1
	return append(slice, vegetable)
}

func FindIndexById(slice []model.Vegetable, id uint) (int, error) {
	for index, vegetable := range slice {
		if(vegetable.Id == id){
			return index, nil
		}
	}
	return 0, fmt.Errorf("Id %d not found", id) 
}

func RemoveByIndex(slice []model.Vegetable, index int) ([]model.Vegetable, error){
	return append(slice[:index], slice[index+1:]...), nil
}

func RemoveById(slice []model.Vegetable, vegetable model.Vegetable) ([]model.Vegetable, error) {
	var index int
	var err error
	index, err = FindIndexById(slice, vegetable.Id)
	if(err != nil){
		return nil, err
	}
	return RemoveByIndex(slice, index)
}

func UpdateByIndex(slice []model.Vegetable, vegetable model.Vegetable, index int) ([]model.Vegetable, error) {
	slice[index].Name = vegetable.Name
	slice[index].Price = vegetable.Price
	return slice, nil
}

func UpdateById(slice []model.Vegetable, vegetable model.Vegetable, id uint) ([]model.Vegetable, error) {
	var index int
	index, _ = FindIndexById(slice, vegetable.Id)
	return UpdateByIndex(slice, vegetable, index)
}