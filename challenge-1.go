package main

import "fmt"

func main () {
	i := 21
	var j bool
	const k float64 = 123.456
	fmt.Println(i) // menampilkan nilai 21
	fmt.Printf("%T\n", i) // menampilkan tipe data i
	fmt.Printf("%c\n",'%') // menampilkan simbol %
	fmt.Println(!j) // menampilkan nilai boolean j = true
	fmt.Printf("%b\n", i) // menampilkan biner dari 21, di gambar compiler ada biner ini
	fmt.Printf("%c\n", 0x042F) // menampilkan unicode rusia Я
	fmt.Printf("%d\n", i) // menampilkan nilai 21/i dengan base 10
	fmt.Printf("%o\n", i) // menampilkan nilai 25 dengan base 8
	fmt.Printf("%x\n", i-6) // menampilkan huruf f dengan base 16
	fmt.Printf("%X\n", i-6) // menampilkan huruf F dengan base 16
	fmt.Printf("%U\n", 'Я') // menampilkan unicode format dari unicode rusia Я
	fmt.Printf("%.6f\n", k) // menampilkan float dengan 6 digit dibelakang koma
	fmt.Printf("%E\n", k) // menampilkan float scientific
}