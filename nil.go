package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}

/**
	Nil :
1. biasanya di dalam bahasa pemrograman lain, object yang belum di inisialisasi maka secara otomatis
	nilainya adalah null atau nil
2. berbeda dengan go-lang, di go-lang saat kita buat variable dengan tipe data tertentu, maka secara
	otomatis akan dibuatkan default valuenya
3. namun di go-lang ada data nil, yaitu data kosong
4. nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice,
	pointer dan channel
*/
