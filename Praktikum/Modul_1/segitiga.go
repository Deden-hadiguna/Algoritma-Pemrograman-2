package main

import "fmt"

func main() {
	var base, height float64

	//input nilai alas dan tinggi dari pengguna
	fmt.Print("Masukkan nilai alas: ")
	fmt.Scan(&base)

	fmt.Print("Masukkan nilai tinggi: ")
	fmt.Scan(&height)

	//Rumus luas segitiga = 1/2 * alas * tinggi
	area := 0.5 * base * height

	//Menampilkan hasil perhitungan luas segitiga
	fmt.Printf("Luas segitiga adalah: %.2f\n", area)
}
