package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Alqo")
	fmt.Println(result)

	fmt.Println(getHello(""))
}

// function return value :
// 1. function bisa mengembalikan data
// 2. untuk memberitahu bahwa function mengembalikan data,
//		kita harus menuliskan tipe data kembalian dari function tersebut
// 3. jika function tersebut kita deklarasikan dengan tipe data pengembalian,
//		maka wajib di dalam functionnya kita harus mengembalikan data
// 4. untuk mengembalikan data dari funtion, kita bisa menggunakan kata kunci return, diikuti dengan datanya
