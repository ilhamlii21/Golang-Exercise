package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() { // main function untuk mengambil data dari struct
	var ilham Customer
	ilham.Name = "Ilham Lii" //ilham.Name untuk mengambil data dari struct
	ilham.Adress = "Jl Dr Cipto"
	ilham.Age = 19

	fmt.Println(ilham)

	lii := Customer{
		Name:   "Lii",
		Adress: "Jl Sarwoendah 3",
		Age:    10,
	}

	fmt.Println(lii)

	naruto := Customer{"Naruto", "Cilacap", 21}
	fmt.Println(naruto)

	naruto.sayHello("Agus")

}
