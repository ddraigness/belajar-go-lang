package main

import (
	"fmt"
)

func main() {
	name := "supremenone"

	switch name {
	case "Alqo":
		fmt.Println("hello alqo")
		fmt.Println("hello alqo")
	case "Mulky":
		fmt.Println("hello mulky")
		fmt.Println("hello mulky")
	default:
		fmt.Println("kenalan dong")
		fmt.Println("kenalan dong")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah oke")
	}
}

// switch Expression :
// 1. selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
// 2. switch expression sangat sederhana dibandingkan if
// 3. biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

// switch dengan short statement :
// sama dengan if, switch juga mendukung short statement sebelum variable yang akan di cek kondisinya

// switch tanpa kondisi :
// 1. kondisi di switch expression tidak wajib
// 2. jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut
//		di setiap casenya.
