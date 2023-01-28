package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Pamor", 28},
		{"Alqo", 26},
		{"Rivan", 27},
		{"Fiqih", 26},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}

/**
	package sort :
1. package sort adalah package yang berisikan utilitas untuk proses pengurutan
2. agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
3. https://golang.org/pkg/sort/
*/
