package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Alqo")
	data.PushBack("Mulky")
	data.PushBack("Aria")
	data.PushFront("Rivan")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)


	// dari depan ke belakang
	// for e := data.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	// dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}

/**
	package container/list :
1. package container/list adalah implementasi struktur data double linked list di go-lang
2. https://golang.org/pkg/container/list/
*/
