package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Change"
	// fmt.Println(slice1)

	// slice1[0] = "Mei Lagi"
	// fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Alqo")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Alqo"
	newSlice[1] = "Mulky"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}

// tipe data slice :
// 1. tipe data slice adalah potongan dari data array
// 2. slice mirip dengan array, yang membedakan adalah ukuran slice bisa berubah
// 3. slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses
//		sebagian atau seluruh data di array

// detail tipe data slice :
// 1. tipe data slice memiliki 3 data, yaitu pointer, length, dan capacity
// 2. pointer adalah penunjuk data pertama di array para slice
// 3. length adalah panjang dari slice, dan
// 4. capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

// membuat slice dari array :
// membuat slice 				|	keterangan
//		array[low:high]					membuat slice dari array dimulai index low sampai index sebelum high
//		array[low:]						membuat slice dari array dimulai index low sampai index akhir di array
//		array[:high]					membuat slice dari array dimulai index 0 sampai index sebelum high
//		array[:]						membuat slice dari array dimulai index 0 sampai index akhir di array
