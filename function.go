package main

import "fmt"

func sayHello() {
	fmt.Println("Helo World")
}

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}

// function :
// 1. sebelumnya kita sudah mengenal sebuah function yang wajib dibuat agar program kita bisa berjalan,
//		yaitu function 'main'.
// 2. function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang.
// 3. cara membuat function sangat sederhana, hanya dengan menggunakan kata kunci 'func' lalu diikuti
//		dengan nama functionnya dan blok kode isi functionnya.
// 4. setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan
// 		kata kunci nama functionnya diikuti tanda kurung buka, kurung tutup
