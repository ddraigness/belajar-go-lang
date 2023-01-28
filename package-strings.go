package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Alqo Mulky", "Alqo"))
	fmt.Println(strings.Split("Alqo Mulky", " "))
	fmt.Println(strings.ToLower("Alqo Mulky"))
	fmt.Println(strings.ToUpper("Alqo Mulky"))
	fmt.Println(strings.Trim("     Alqo Mulky    ", " "))
	fmt.Println(strings.ReplaceAll("Alqo Alqo Alqo Alqo", "Alqo", "Rivan"))
}

/**
	package strings :
1. package strings adalah package yang berisikan function-function untuk memanipulasi
	tipe data String
2. ada banyak sekali function yang bisa kita gunakan
3. https://golang.org/pkg/strings/
*/

/**
	beberapa function di package strings :
Function :							| Kegunaan :
1. strings.Trim(string, cutset)			Memotong cutset di awal dan akhir string
2. strings.ToLower(string)				Membuat semua karakter string menjadi lower case
3. strings.ToUpper(string)				Membuat semua karakter string menjadi upper case
4. strings.Split(string, separator)		Memotong string berdasarkan separator
5. strings.Contains(string, search)		Mengecek apakah string mengandung string lain
6. strings.ReplaceAll(string, from, to)	Mengubah semua string dari from ke to
*/
