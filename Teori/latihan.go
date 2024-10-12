package main

import "fmt"

func main() {
	var jari, tinggi float32
	var volume float32
	fmt.Scan(&jari, &tinggi)
	volume = volumeTabung(jari, tinggi)
	fmt.Println(volume)
}

func volumeTabung(r, t float32) float32 {
	var hasil float32
	hasil = 3.14 * r * r * t
	return hasil
}
