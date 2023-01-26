package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke", i)
		if i == 5 {
			break
		}
	}

}

// Break & Continue :
// 1. break & continue adalah kata kunci yang bisa digunakan dalam perulangan
// 2. break digunakan untuk menghentikan seluruh perulangan
// 3. continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung
//		melanjutkan ke perulangan selanjutnya
