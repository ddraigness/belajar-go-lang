package main

import (
	"reflect"
	"fmt"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{"Alqo"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")
	max := structField.Tag.Get("max")

	fmt.Println(structField.Name)
	fmt.Println(sampleType.NumField())
	fmt.Println(required)
	fmt.Println(max)

	fmt.Println(IsValid(sample))
}

/**
	package reflect :
1. dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode
	kita pada saat aplikasi sedang berjalan
2. hal ini bisa dilakukan di go-lang dengan menggunakan package reflect
3. fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, anda bisa
	eksplorasi package reflect ini secara otodidak
4. Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
5. https://golang.org/pkg/reflect/
*/
