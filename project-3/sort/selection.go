package sort

import (
	"project-2/model"
)

func Selection(slice []model.Vegetable) []model.Vegetable{
	for i := 0; i < len(slice) -1; i++ {
		// assume lowest value at index i
		minIndex := i

		// find index with minimum item
		for j := i + 1; j < len(slice); j++ {
			if (slice[j].Id < slice[minIndex].Id){
				minIndex = j
			}
		}

		// swap item with minimum index with i
		slice[i], slice[minIndex] = slice[minIndex], slice[i]
	}
	return slice
}