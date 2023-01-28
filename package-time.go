package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2023, 01, 31, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2022-12-31T15:04:05Z")
	fmt.Println(parse)
}

/**
	package time :
1. package time adalah package yang berisikan fungsionalitas untuk management waktu di go-lang
*/

/**
	beberapa function di package time :
Function :					| Kegunaan :
1. time.Now()					untuk mendapatkan waktu saat ini
2. time.Date(...)				untuk membuat waktu
3. time.Parse(layout, string)	untuk memparsing/menguraikan waktu dari string
*/
