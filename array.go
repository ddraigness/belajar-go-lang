package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Alqo"
	names[1] = "Mulky"
	names[2] = "Aria"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))
}

// tipe data array :
// 1. array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
// 2. saat membuat array, kita perlu mennetukan jumlah data yang bisa ditampung oleh array tersebut
// 3. daya tampung array tidak bisa bertambah setelah array dibuat

// index di array :
// index								|	data
//		0										Alqo
//		1										Mulky
//		2										Aria
