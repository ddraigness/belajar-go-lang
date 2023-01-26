package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello ", nameFiltered)
// }

func spamFilter(name string) string {
	if name == "Babone" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Alqo", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Babone", filter)
}

// function as parameter :
// 1. function tidak hanya bisa kita simpan di dalam variable sebagai value
// 2. namun juga bisa kita gunakan sebagai parameter untuk function lain

// function type declarations :
// 1. kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
// 2. type declarations juga bisa digunakan untuk membuat alias function, sehingga akan
//		mempermudah kita menggunakan function sebagai parameter
