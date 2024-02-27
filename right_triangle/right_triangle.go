package main

import "fmt"

/*
- for i := 0; i < 5; i++ => pertama dibuatkan barisnya (rows), karena menggunakan prinln, akan membuat baris baru setiap perulangan
- j := 0; j <= i; j++ => lalu dibuatkan data berapa banyak bintang tiap baris, karena menggunakan print, maka tidak membuat baris baru
- kenapa j <= i => agar perulangan baris bertahap mengikuti perulangan baris, jadi baris pertama akan membuat 1 bintang, baris kedua 2 bintang dst
*/

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
