package main

import "fmt"

func main() {
	var (
		name1       = "Alqo"
		name2       = "Subur"
		result bool = name1 == name2
	)
	fmt.Println(result)

	var (
		value1 = 100
		value2 = 200
	)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}

// operasi perbandingan :
// 1. operasi perbandingan adalah operasi untuk membandingkan dua buah data
// 2. operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
// 3. jika hasil operasinya adalah benar, maka nilainya true
// 4. jika hasil operasinya adalah salah, maka nilainya adalah false
