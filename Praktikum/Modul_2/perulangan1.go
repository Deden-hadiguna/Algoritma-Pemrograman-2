package main
import (
	"fmt"
)
func main (){
	var angka int
	total := 0
	fmt.Println("Masukkan angka positif untuk dijumlahkan. Masukkan angka negatif untuk keluar.")
	//simulasi while loop
	for {
		fmt.Println("Masukkan angka: ")
		fmt,Scanln(&angka)
		if angka < 0 { //Keluar dari loop jika angka negatif
			break
		}
		total += angka //Tambahkan angka ke total
	}
	fmt.Printf("Total jumlah dari angka yang dimasukkan adalah: %d\n", total)
}