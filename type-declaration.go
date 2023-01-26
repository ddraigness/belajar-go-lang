package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noKtpAlqo noKTP = "3724678236478236"
	var marriedStatus Married = true
	fmt.Println(noKtpAlqo)
	fmt.Println(marriedStatus)
}

// type declarations :
// 1. type declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
// 2. type declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada,
// 		dengan tujuan agar lebih mudah dimengerti
