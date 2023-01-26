package main

import "fmt"

func main() {
	var name string

	name = "Alqo Mulky"
	fmt.Println(name)

	name = "Mulky Aria"
	fmt.Println(name)

	var friendName = "Irgi"
	fmt.Println(friendName)

	var age = 26
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstname  = "Alqo"
		middlename = "Mulky"
		lastname   = "Aria"
	)
	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
}

// variable :
// 1. variable adalah tempat untuk menyimpan data.
// 2. variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau.
// 3. di go-lang variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data
//		yang berbeda-beda jenis, kita harus membuat beberapa variable.
// 4. untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable
// 		dan tipe datanya.

// deklarasi multiple variable :
// 1. di go-lang kita bisa membuat variable secara sekaligus banyak
// 2. code yang dibuat akan lebih bagus dan mudah dibaca
