package main

import "fmt"

func getFullName() (string, string, string) {
	return "Alqo", "Mulky", "Aria"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

// returning multiple value :
// 1. function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
// 2. untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe
//		data return valuenya di function

// menghiraukan return value :
// 1. multiple return value wajib ditangkap semua valuenya
// 2. jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda _ (garis bawah/underscore)
