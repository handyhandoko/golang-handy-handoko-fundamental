package handler

import (
	"fmt"
	"project-2/service"
)

type VegetableHandler struct {
	Service service.VegetableService
}

func (handler *VegetableHandler) GetAllVegetables(){
	vegetables, _ := handler.Service.GetAllVegetables()

	for index, vegetable := range vegetables {
		fmt.Printf("%d. %s\n", index+1, vegetable)
	}
}