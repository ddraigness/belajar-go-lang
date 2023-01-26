package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye

	result := goodbye("Aria")
	fmt.Println(goodbye("Alqo"))
	fmt.Println(result)
}

// function value :
// 1. function adalah first class citizen
// 2. function juga merupakan tipe data, dan bisa disimpan di dalam variable
