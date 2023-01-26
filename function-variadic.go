package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 90, 20, 30, 40, 50)
	fmt.Println(total)

	slice := []int{10, 20, 30, 40, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}

// variadic function :
// 1. parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
// 2. varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam array
// 3. apa bedanya dengan parameter biasa dengan tipe data array?
//		- jika parameter tipe array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function.
//		- jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu,
//			cukup gunakan tanda koma

// slice parameter :
// 1. kadang ada kasus dimana kita menggunakan variadic function, namun memiliki variable berupa slice.
// 2. kita bisa menjadi slice sebagai varargs parameter
