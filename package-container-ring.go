package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)
	for i := 0; 1 < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}

/**
	package container/ring :
1. package container/ring adalah implementasi struktur data circular list
2. circular list adalah struktur data ring, dimana diakhir element akan kembali ke
	element awal (HEAD)
3. https://golang.org/pkg/container/ring/
*/
