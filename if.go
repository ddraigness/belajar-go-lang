package main

import "fmt"

func main() {
	var name = "Aria"

	if name == "Alqo" {
		fmt.Println("Hello Alqo")
	} else if name == "Aria" {
		fmt.Println("Halo Aria")
	} else if name == "Surya" {
		fmt.Println("Hallo Surya")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama anda sudah benar")
	}
}

// if Expression :
// 1. if adalah salah satu kata kunci yang digunakan untuk percabangan
// 2. percabangan artinya kita bisa mengeksekusi kode program tertenru ketika suatu kondisi terpenuhi
// 3. hampir di semua bahasa pemrograman mendukung if expression

// else Expression :
// 1. blok if akan dieksekusi ketika kondisi if bernilai true
// 2. kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
// 3. hal ini bisa dilakukan menggunakan else Expression.

// else if Expression :
// 1. kadang dalam if, kita butuh membuat beberapa kondisi
// 2. kasus seperti ini, kita bisa menggunakan Else If expression

// if dengan short statement :
// 1. if mendukung short statement sebelum kondisi
// 2. hal ini sangan cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi
