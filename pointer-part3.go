package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	// fmt.Println("Di Method", man.Name)
}

func main() {
	alqo := Man{"Alqo"}
	alqo.Married()

	fmt.Println(alqo.Name)
}

/**
	pointer di method :
1. walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses
	di dalam method adalah pass by value
2. sangat di rekomendasikan menggunakan pointer di method, sehingga tidak boros memory
	karena harus selalu diduplikasi ketika memanggil method
*/
