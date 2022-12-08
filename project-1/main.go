package main

import (
	"fmt"
)

func main(){
	var vegetables = []string{"sawi", "bayam", "kangkung", "kol", "pare"}
	var selectedMenu int
	for selectedMenu < 3{
		selectedMenu = printMainMenu()
		switch(selectedMenu){
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

func printMainMenu() int{
	fmt.Println("Selamat datang di OndeMaret. Selamat belanja.")
	fmt.Println("Masukkan (1) untuk menambah data.")
	fmt.Println("Masukkan (2) untuk melihat semua data.")
	fmt.Println("Masukkan (3) keluar dari aplikasi.")

	var selectedMenu int
	fmt.Scanln(&selectedMenu)
	return selectedMenu
}

func showData(index int, vegetableName string){
	fmt.Println(vegetableName)
}

func listData(vegetables []string){
	for index, vegetable := range(vegetables){
		fmt.Printf("%d. %s\n", index+1 , vegetable)
	}
	fmt.Println(`Untuk melihat, mengupdate, atau menghapus data.
		Silahkan masukkan nomor yang tertera di sisi kiri nama barang.
		Untuk kembali ke menu utama, masukkan angka 0.`)

	var selectedMenu int
	fmt.Scanln(&selectedMenu)
	if index:= selectedMenu-1; (selectedMenu == 0){
		main()
	} else if(selectedMenu <= len(vegetables)){
		showData(index, vegetables[index])
	// } else {
	// 	fmt.Printf("Invalid input. masukkan angka lebih kecil dari %d", len(vegetables))
	}
}

func addData(vegetables []string) []string{
	fmt.Println("Masukkan nama sayur:")
	var vegetableName string
	fmt.Scanln(&vegetableName)
	return append(vegetables, vegetableName)
}