package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Alqo"
	sayHelloTo(firstName, "Mulky")
	sayHelloTo("Surya", "Subur")
}

// function parameter :
// 1. saat membuat function, kadang-kadang kita membutuhkan data dari luar, atau kita sebut parameter.
// 2. kita bisa menambahkan parameter di function, bisa lebih dari satu.
// 3. parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya
//		yang sudah kita buat
// 4. namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut,
//		kita wajib memasukkan data ke parameternya.
