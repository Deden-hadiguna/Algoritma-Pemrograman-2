package main

import "fmt"

// fungsi untuk menghitung panjang dan lebar dari luas
func hitungPanjangLebar(luas float64, rasio float64) (float64, float64) {
	// Misalkan lebar = rasio * panjang
	lebar := luas / rasio
	panjang := luas / lebar
	return panjang, lebar
}

// fungsi untuk menghitung keliling persegi panjang
func hitungKeliling(panjang float64, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func main() {
	// mendefinisikan luas dan rasio
	var luas, rasio float64

	fmt.Print("Masukkan luas persegi panjang: ")
	fmt.Scan(&luas)

	fmt.Print("Masukkan rasio lebar terhadap panjang (misal 0.5 untuk lebar setengah dari panjang): ")
	fmt.Scan(&rasio)

	// menghitung panjang dan lebar
	panjang, lebar := hitungPanjangLebar(luas, rasio)

	// menampilkan hasil
	fmt.Printf("Panjang: %.2f, Lebar: %.2f\n", panjang, lebar)

	// menghitung dan menampilkan keliling
	keliling := hitungKeliling(panjang, lebar)
	fmt.Printf("Keliling persegi panjang: %.2f\n", keliling)
}
