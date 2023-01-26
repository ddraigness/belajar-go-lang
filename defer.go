package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(10)
}

// defer :
// 1. defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function
//		selesai di eksekusi
// 2. defer function akan selalu dieksekusi walaupun terjadi error di function yang di eksekusi
