package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}

/**
	pointer di function :
1. saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan
	di copy lalu dikirim ke function tersebut
2. oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan
	pernah berubah
3. hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
4. namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
5. untuk melakukan ini, kita juga bisa menggunakan pointer di function
6. untuk menjadikan sebuah parameter menjadi pointer, kita bisa menggunakan operator '*' di parameternya
*/
