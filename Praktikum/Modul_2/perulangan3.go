package main

import (
	"fmt"
)

func main() {
	maxNumber := 10 //batas maksimum
	number := 1     //angka awal
	//simulasi repeat-until
	for finished := false; !finished; {
		fmt.Println("angka:", number) //cetak angka saat ini
		//cetak kondisi keluar
		if number >= maxNumber {
			finished = true //ubah kondisi untuk keluar dari loop
		}
		number++ //increment angka
	}
}
