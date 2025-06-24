package main

import "fmt"

func sayHelloto(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloto("ilham", "Lii")
	sayHelloto("Naruto", "Uzumaki")

}
