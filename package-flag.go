package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "admin", "Put your database username")
	password := flag.String("password", "P@ssw0rd", "Put your database password")

	flag.Parse()

	fmt.Println("Hostname :", *host)
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
}

/**
	package flag :
1. package flag berisikan fungsionalitas untuk memparsing command line argument
2. https://golang.org/pkg/flag/
*/
