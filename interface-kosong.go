package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(3)
	fmt.Println(data)
}

/**
	Interface Kosong :
1. go-lang bukanlah bahasa pemrograman yang berorientasi objek
2. biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa
	dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
3. contoh di Java ada java.lang.Object
4. untuk menangani kasus seperti ini, di go-lang kita bisa menggunakan interface kosong
5. interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat
	secara otomatis semua tipe data akan menjadi implementasinya
*/

/**
	penggunaan interface kosong di go-lang :
1. ada banyak contoh penggunaan interface kosong di go-lang, seperti :
	a. fmt.Println(a ...interface{})
	b. panic(v interface{})
	c. recover{} interface{}
	d. dan lain-lain
*/
