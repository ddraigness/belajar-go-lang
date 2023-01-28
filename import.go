package main

import (
	"belajar-go-lang/helper"
	"fmt"
)

func main() {
	helper.SayHello("Alqo")
	// helper.Application("Alqo") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}

/**
	import :
1. secara standar, file go-lang hanya bisa mengakses file go-lang lainnya yang berada
	dalam package yang sama
2. jika kita ingin mengakses file go-lang yang berada diluar package, maka kita bisa
	menggunakan import
*/
