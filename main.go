package main

import "fmt"

func main() {

	// nilai := 80

	// if nilai >= 90 {
	// 	fmt.Println("Sangat Memuaskan")
	// } else {
	// 	fmt.Println("Nilai dibawah 90")
	// }

	// if nilai >= 90 && nilai <= 100 {
	// 	fmt.Println("Sangat Memuaskan")
	// } else if nilai >= 80 && nilai < 90 {
	// 	fmt.Println("Memuaskan")
	// } else if nilai >= 70 && nilai  < 80 {
	// 	fmt.Println("Baik")
	// } else if nilai >= 60 && nilai < 70 {
	// 	fmt.Println("Cukup")
	// } else if nilai < 60 {
	// 	fmt.Println("Tidak Lulus")
	// } else {
	// 	fmt.Println("Nilai Tidak Diketahui")
	// }

	angka := 1
	switch angka {
	case 1:
		{
			fmt.Println("Angka 1")
			break
		}
	case 2:
		{
			fmt.Println("Angka 2")
			break
		}
	default:
		{
			fmt.Println("Angka tidak diketahui")
		}
	}

	nilai := 90

	switch {
	case nilai >= 90 && nilai <= 100:
		fmt.Println("Sangat Memuaskan")
		break
	case nilai >= 80 && nilai < 90:
		fmt.Println("Memuaskan")
		break 	
	}
}