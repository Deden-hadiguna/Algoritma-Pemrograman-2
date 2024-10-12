package main

import "fmt"

//fungsi untuk menghitung luas persegi panjang
func hitungLuas(panjang float64, lebar float64) float64 {
	return panjang * lebar
}

func main() {
	//mendefinisikan panjang dan lebar
	var panjang, lebar float64

	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&panjang)

	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)

	//memanggil fungsi hitungLuas dan menampilkan hasilnya
	luas := hitungLuas(panjang, lebar)
	fmt.Printf("Luas persegi panjang: %.2f\n", luas)
}
