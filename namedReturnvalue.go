package main

import "fmt"

func getCompleteName() (firstname, middleName, lastName string) { // penggunaan tipedata jika sama bisa dibelakang
	firstname = "Ilham"
	middleName = "Lii"
	lastName = "Assidaq"

	return firstname, middleName, lastName

}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
