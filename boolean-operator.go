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

// operasi boolean :
//	operator :									| 	keterangan :
// 		&&												Dan
//		||												Atau
//		!												kebalikan

// operasi && :
//	nilai 1			|	operator				| 	nilai 2					| 	hasil
//		true				&&							true						true
//		true				&&							false						false
//		false				&&							true						false
//		false				&&							false						false

// operasi || :
//	nilai 1			|	operator				| 	nilai 2					| 	hasil
//		true				||							true						true
//		true				||							false						true
//		false				||							true						true
//		false				||							false						false
