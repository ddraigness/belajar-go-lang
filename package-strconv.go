package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)
}

/**
	package strconv :
1. sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int64
2. bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? misal,
	dari int ke string, atau sebaliknya
3. hal tersebut bisa kita lakukan dengan bantuan package strconv(string conversion)
4. https://golang.org/pkg/strconv/
*/

/**
	beberapa function di package strconv :
function :						| kegunaan :
1. strconv.parseBool(string)		Mengubah string ke bool
2. strconv.parseFloat(string)		Mengubah string ke float
3. strconv.parseInt(string)			Mengubah string ke int64
4. strconv.FormatBool(bool)			Mengubah bool ke string
5. strconv.FormatFloat(float,...)	Mengubah float64 ke string
6. strconv.FormatInt(int,...)		Mengubah int64 ke string
*/
