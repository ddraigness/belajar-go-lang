package main

import (
	"belajar-go-lang/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}

/**
	package initialization :
1. saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika
	package kita diakses
2. ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk
	berkomuikasi dengan database, kita membuat function inisialisasi untuk membuka
	koneksi ke database
3. untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup
	membuat function dengan nama init
*/

/**
	Blank Identifier :
1. kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi
	salah satu function yang ada di package
2. secara default, go-lang akan komplen ketika ada package yang di import namun tidak digunakan
3. untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package
	ketika melakukan import
*/
