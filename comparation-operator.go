package main

import "fmt"

func main() {
	var (
		name1       = "Alqo"
		name2       = "Subur"
		result bool = name1 == name2
	)
	fmt.Println(result)

	var (
		value1 = 100
		value2 = 200
	)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
