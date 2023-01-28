package helper

import "fmt"

var version = 1
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello" + name)
}

func sayGoodBye(name string) string {
	fmt.Println("Goodbye", name)
}

/**
	Access Modifier :
1. di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk
	menentukan access modifier terhadap suatu function atau variable
2. di go-lang, untuk menentukan access modifier, cukup dengan nama function atau variable
3. jika namanya di awali dengan huruf besar, maka artinya bisa di akses dari package
	lain, jika dimulai dengan huruf kecil, artinya tidak bisa di akses dari package lain
*/
