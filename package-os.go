package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
	// fmt.Println(args[1])
	// fmt.Println(args[2])
	// fmt.Println(args[3])
}

/**
	package os :
1. go-lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
2. package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen
	(bisa digunakan disemua sistem operasi)
3. https://golang.org/pkg/os/
*/
