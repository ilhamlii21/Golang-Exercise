package main

import "fmt"

type HasName interface {
	GetName() string //getname adalah method return string
}

/*

type Hasname interface berfungsi untuk
mengambil data dari struct melalui method GetName

*/

func SayHello(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

/*

sayhello memiliki method return string yaitu
"hello" dan mengambil data dari struct melalui method GetName

*/

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name //person dibuat untuk mengambil data dari struct
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Ilham Lii Assidaq"}
	SayHello(person)

	animal := Animal{Name: "Toothles"}
	SayHello(animal)
}
