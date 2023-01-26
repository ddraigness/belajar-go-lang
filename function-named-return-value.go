package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Pamor"
	middleName = "Serenada"
	lastName = "Babone"

	return

}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// named return values:
// 1. biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value,
//		maka kita hanya mendeklarasikan tipe data return value di function
// 2. namun kita juga bisa membuat variable secara langsung di tipe data return functionnya
