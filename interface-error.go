package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

	// contohError := errors.New("Ups Error")
	// fmt.Println(contohError.Error())
}

/**
error interface :
1. go-lang memiliki interface yang digunakan sebagai kontrak untuk membuat error,
nama interfacenya adalah error.

type error interface {
	Error() string
}
*/

/**
	membuat error :
1. untuk membuat error, kita tidak perlu manual
2. go-lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat
	di package errors (Package akan kita bahas secara detail di materi tersendiri)
*/
