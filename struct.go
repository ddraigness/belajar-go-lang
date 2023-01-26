package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHai() {
	fmt.Println("Hai bro", a.Name)
}

func main() {
	var alqo Customer
	alqo.Name = "Alqo Mulky Aria"
	alqo.Address = "Indonesia/Jakarta"
	alqo.Age = 26

	alqo.sayHi("Randa")
	alqo.sayHai()

	// fmt.Println(alqo)
	// fmt.Println(alqo.Name)
	// fmt.Println(alqo.Address)
	// fmt.Println(alqo.Age)

	// babone := Customer{
	// 	Name:    "Randa",
	// 	Address: "Cakung",
	// 	Age:     26,
	// }
	// fmt.Println(babone)

	// irgi := Customer{"Irgi", "Kemayoran", 21}
	// fmt.Println(irgi)
}

/**
	struct :
1. struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih
	tipe data lainnya dalam satu kesatuan
2. struct biasanya representasi data dalam program aplikasi yang kita buat
3. data di struct disimpan dalam field
4. sederhananya struct adalah kumpulan dari field
*/

/**
	membuat data struct :
1. struct adalah template data atau prototype data
2. struct tidak bisa langsung digunakan
3. namun kita bisa membuat data/object dari struct yang telah kita buat
*/

/**
	struct literals :
sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa
kita gunakan untuk membuat data dari struct
*/

/**
	struct method :
1. struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk
	function
2. namun jika kita ingin menambahkan method ke dalam struct, sehingga seakan-akan sebuah struct
	memiliki function
3. method adalah function
*/
