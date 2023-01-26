package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message :", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Alqo")
}

// panic :
// 1. panic function adalah function yang bisa kita gunakan untuk menghentikan program
// 2. panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
// 3. saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

// recover :
// 1. recover adalah function yang bisa kita gunakan untuk menangkap data panic
// 2. dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
