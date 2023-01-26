package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }
// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("root", func(name string) bool {
		return name == "admin"
	})
}

// anonymous function :
// 1. sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut.
// 2. namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter
//		tanpa harus membuat function terlebih dahulu
// 3. hal tersebut dinamakan anonymous function, atau function tanpa nama
