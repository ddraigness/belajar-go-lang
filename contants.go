package main

import "fmt"

func main() {
	const (
		firstName string = "Alqo"
		lastName         = "Mulky"
		value            = 1000
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}

// constant :
// 1. constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
// 2. cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan
// 		adalah const, bukan var
// 3. saat pembuatan constant, kita wajib langsung menginisialisasikan datanya

// deklarasi multiple constant :
// 1. sama seperti variable, di go-lang juga kita bisa membuat constant secara sekaligus banyak.
