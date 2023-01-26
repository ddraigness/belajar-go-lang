package main

import "fmt"

func main() {
	name := "Alqo"
	counter := 0

	increment := func() {
		name := "Aria"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}

// closures :
// 1. closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya
//		dalam scope yang sama
// 2. harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi
