package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}

	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)
}

/**
	type assertions :
1. type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
2. fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
*/

/**
	type assertions menggunakan switch :
1. saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
2. jika panic dan tidak ter-recover, maka otomatis program kita akan mati
3. agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
*/
