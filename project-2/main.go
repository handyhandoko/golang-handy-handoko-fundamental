package main

import (
	"fmt"
	"project-2/service"
	"project-2/domain"
	"project-2/handler"
)

var vegetables = []string{}

func main() {
	var selectedMenu int
	for selectedMenu < 3 {
		selectedMenu = printMainMenu()
		switch selectedMenu {
		case 1:
			vegetables = addData(vegetables)
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

func removeByIndex(slice []string, index int) {
	vegetables = append(vegetables[:index], vegetables[index+1:]...)
}

func updateData(slice []string, index int) {
	fmt.Println("Masukkan nama baru.")
	var newVegetableName string
	fmt.Scanln(&newVegetableName)
	vegetables[index] = newVegetableName
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
		removeByIndex(vegetables, index)
	}
}

func listData(vegetables []string) {
	vegetableRepository := domain.NewVegetableRepositoryStub()
	vegetableService := service.NewVegetableService(vegetableRepository)
	vegetableHandler := handler.VegetableHandler{vegetableService}
	vegetableHandler.GetAllVegetables()

// 	fmt.Println(`Untuk melihat, mengupdate, atau menghapus data.
// Silahkan masukkan nomor yang tertera di sisi kiri nama barang.
// Untuk kembali ke menu utama, masukkan angka 0.`)

// 	var selectedMenu int = len(vegetables) + 1
// 	for selectedMenu > len(vegetables) {
// 		fmt.Scanln(&selectedMenu)
// 		if index := selectedMenu - 1; selectedMenu == 0 {
// 			break
// 		} else if selectedMenu <= len(vegetables) {
// 			showData(index)
// 		} else {
// 			fmt.Printf("Invalid input. masukkan angka lebih kecil dari %d", len(vegetables) + 1)
// 		}
// 	}
}

func addData(vegetables []string) []string {
	fmt.Println("Masukkan nama sayur:")
	var vegetableName string
	fmt.Scanln(&vegetableName)
	return append(vegetables, vegetableName)
}