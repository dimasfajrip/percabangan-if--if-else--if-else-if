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

	// angka := 1
	// switch angka {
	// case 1:
	// 	{
	// 		fmt.Println("Angka 1")
	// 		break
	// 	}
	// case 2:
	// 	{
	// 		fmt.Println("Angka 2")
	// 		break
	// 	}
	// default:
	// 	{
	// 		fmt.Println("Angka tidak diketahui")
	// 	}
	// }

	// nilai := 90

	// switch {
	// case nilai >= 90 && nilai <= 100:
	// 	fmt.Println("Sangat Memuaskan")
	// 	break
	// case nilai >= 80 && nilai < 90:
	// 	fmt.Println("Memuaskan")
	// 	break 	
	// }

	// i := 1
	// for i <= 10 {
	// 	fmt.Printf("Tentangkode ke %v\n", i)
	// 	i++
	// }

	// for i:= 1; i<=10; i+=2{
	// 	fmt.Printf("Tentangkode ke %v \n", i)
	// }

	// for i:= 1; i<=10; i--{
	// 	fmt.Printf("Tentangkode ke %v \n", i)
	// }

	// break => menghentikan perulangan setelah memenuhi suatu kondisi tertentu
	// i:= 1
	// for i <= 10 {

	// 	if i == 5 {
	// 		fmt.Println("Perulangan sudah mencapai 5, proses perulangan di stop")
	// 		break
	// 	}

	// 	fmt.Printf("Belajar Golang ke %v \n", i)

	// 	i++
	// }

	// fmt.Println("Baris ini tereksekusi setelah for")

	//continue => melanjutkan perulangan setelah memenuhi suatu kondisi tertentu
	// i:= 1
	// for i <= 10 {

	// 	if i == 5 {
	// 		fmt.Println("Perulangan sudah mencapai 5, proses perulangan akan loncat ke yg 7")
	// 		i += 2
	// 		//  i = 7
	// 		continue
	// 	}

	// 	fmt.Printf("Belajar Golang ke %v \n", i)

	// 	i++
	// }

	// fmt.Println("Baris ini tereksekusi setelah for")


	//goto => digunakan untuk keluar dari proses perulangan, dan loncat ke baris  tertentu (label goto)
	i:= 1
	for i <= 10 {

		if i == 5 {
			goto loncat_kesini
		}

		fmt.Printf("Belajar Golang ke %v \n", i)

		i++
	}

	loncat_kesini:
	fmt.Println("Jangan lupa subscribe")
}