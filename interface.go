package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Alqo"}
	animal := Animal{Name: "Kucing"}
	SayHello(person)
	SayHello(animal)
}

/**
	Interface :
1. Interface adalah tipe data abstract, dia tidak memiliki implementasi lansung
2. sebuah interface berisikan definisi-definisi method
3. biasanya interface digunakan sebagai kontrak
*/

/**
	Implementasi Interface :
1. setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
2. sehingga kita tidak perlu mengimplementasikan interface secara manual
3. hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan
	secara eksplisit akan menggunakan interface mana
*/
