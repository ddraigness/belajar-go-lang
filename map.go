package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Alqo",
		"address": "Jakarta",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Alqo"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}

// tipe data map :
// 1. pada array atau slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
// 2. Map adalah tipe data lain yang berisikan kumpulan data yang sama,
//		namun kita bisa menentukan jenis tipe data index yang akan kita gunakan.
// 3. sederhananya, map ada tipe data kumpulan key-value (kata kunci - nilai),
//		dimana kata kuncinya bersifat unik, tidak boleh sama.
// 4. berbeda dengan array dan slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya,
//		asalkan kata kuncinya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis
//		data sebelumnya akan diganti dengan data baru.

// function map :
// Operasi 							|	keterangan
//		len(map)							Untuk mendapatkan jumlah data di map
//		map[key]							Mengambil data di map dengan key
//		map[key] = value					Mengubah data di map dengan key
//		make(map[TypeKey]TypeValue)			Membuat map baru
//		delete(map, key)					Menghapus data di map dengan key
