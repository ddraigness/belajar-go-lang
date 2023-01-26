package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1
	// var address1 *Address = &Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}

	// 	variablenya bisa juga menggunakan yang dibawah ini :
	// address1 := Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	// address2 := &address1
	// address3 := &address1
	// address4 := &Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}

	// address2 := &address1 // operator &

	address2.City = "Jakarta Timur"

	address2 = &Address{"Bandung", "Jawa Barat", "Indonesia"}
	*address3 = Address{"Bandung", "Jawa Barat", "Indonesia"} // operator '*'

	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)
}

/**
	Pass by Value :
1. secara default di go-lang semua variable itu di passing by value, bukan by reference
2. artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
	sebenarnya yang dikirim adalah duplikasi valuenya
*/

/**
	Pointer :
1. pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa
	menduplikasi data yang sudah ada
2. sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
*/

/**
	Operator & :
1. untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa
	menggunakan operator '&' diikuti dengan nama variable nya
*/

/**
	Operator * :
1. saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
2. semua variable yang mengacu ke data yang sama tidak akan berubah
3. jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa
	menggunakan operator '*'
*/

/**
	function new :
1. sebelumnya untuk membuat pointer dengan menggunakan operator '&'
2. go-lang juga memiliki function new yang bisa digunakan untuk membuat pointer
3. namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
*/
