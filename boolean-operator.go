package main

import "fmt"

func main() {
	var (
		ujian   = 80
		absensi = 75
	)

	var (
		lulusUjian   = ujian >= 80
		lulusAbsensi = absensi >= 80

		lulus = lulusAbsensi && lulusUjian
	)
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)
}
