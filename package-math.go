package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))
	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Ceil(1.3))

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}

/**
	package math :
1. package math merupakan package yang berisikan constant dan fungsi matematika
2. https://golang.org/pkg/math
*/

/**
	beberapa function di package math :
	Fungsi :					| Kegunaan :
1. math.Round(float64)				Membulatkan float64 keatas atau kebawah, seusai dengan yang paling dekat
2. math.Floor(float64)				Membulatkan float64 kebawah
3. math.Cell(float64)				Membulatkan float64 keatas
4. math.Max(float64, float64)		Mengembalikan nilai float64 paling besar
5. math.Min(float64, float64)		Mengembalikan nilai float64 paling kecil
*/
