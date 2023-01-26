package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Alqo", "Mulky", "Aria"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Alqo"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}

// For :
// 1. dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
// 2. salah satu fitur perulangan adalah for loops

// For dengan statement:
// 1. dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan di for
// 2. Init statement, yaitu statement sebelum for di eksekusi
// 3. Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan

// For Range :
// 1. for bisa digunakan untuk melakukan iterasi terhadap semua data collection
// 2. data collection contohnya array, slice dan map
