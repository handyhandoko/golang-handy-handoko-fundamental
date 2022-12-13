package main

import (
	"fmt"
	"project-2/repository"
	"project-2/model"
)

var vegetables = []model.Vegetable{
	model.Vegetable {1, "sawi", 4500},
	model.Vegetable {2, "bayam", 2000},
	model.Vegetable {3, "kangkung", 1000},
	model.Vegetable {4, "kol", 5000},
	model.Vegetable {5, "pare", 3000},
}

func main() {
	var selectedMenu int
	for selectedMenu < 3 {
		selectedMenu = printMainMenu()
		switch selectedMenu {
		case 1:
			addData(vegetables)
		case 2:
			listData(vegetables)
		case 3:
			continue
		default:
			fmt.Println("Invalid input. Silahkan masukkan angka 1, 2 atau 3.")
		}
	}
}

func printMainMenu() int {
	fmt.Println("Selamat datang di OndeMaret. Selamat belanja.")
	fmt.Println("Masukkan (1) untuk menambah data.")
	fmt.Println("Masukkan (2) untuk melihat semua data.")
	fmt.Println("Masukkan (3) keluar dari aplikasi.")

	var selectedMenu int
	fmt.Scanln(&selectedMenu)
	return selectedMenu
}

func inputData() model.Vegetable {
	fmt.Println("Masukkan nama:")
	var newVegetableName string
	fmt.Scanln(&newVegetableName)
	fmt.Println("Masukkan harga: ")
	var newVegetablePrice uint
	fmt.Scanln(&newVegetablePrice)
	return model.Vegetable { 0, newVegetableName, newVegetablePrice}
}

func updateData(slice []model.Vegetable, index int) {
	var vegetableUpdate model.Vegetable
	vegetableUpdate = inputData()
	vegetables, _ = repository.UpdateByIndex(vegetables, vegetableUpdate, index)
}

func showData(index int) {
	fmt.Println(vegetables[index])
	fmt.Println(`Untuk mengupdate data, tekan 1.
Untuk menghapus data, tekan 2.
Masukkan selain 1 atau 2 untuk kembali ke menu utama.`)

	var selectedMenu int
	fmt.Scanln(&selectedMenu)
	if selectedMenu == 1 {
		updateData(vegetables, index)
	} else if selectedMenu == 2 {
		repository.RemoveByIndex(vegetables, index)
	}
}

func listData(vegetables []model.Vegetable) {
	for index, vegetable := range vegetables {
		fmt.Printf("%d. %s\n", index+1, vegetable)
	}
	fmt.Println(`Untuk melihat, mengupdate, atau menghapus data.
Silahkan masukkan nomor yang tertera di sisi kiri nama barang.
Untuk kembali ke menu utama, masukkan angka 0.`)

	var selectedMenu int = len(vegetables) + 1
	for selectedMenu > len(vegetables) {
		fmt.Scanln(&selectedMenu)
		if index := selectedMenu - 1; selectedMenu == 0 {
			break
		} else if selectedMenu <= len(vegetables) {
			showData(index)
		} else {
			fmt.Printf("Invalid input. masukkan angka lebih kecil dari %d", len(vegetables) + 1)
		}
	}
}

func addData(vegetables []model.Vegetable) {
	var vegetableNew model.Vegetable
	vegetableNew = inputData()
	vegetables = repository.AddItem(vegetables, vegetableNew)
}