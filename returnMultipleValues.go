package main

import "fmt"

func getFullName() (string, string) {
	return "Ilham", "Lii"
}

// func main() {
// 	firstName, lastName := getFullName()
// 	fmt.Println(firstName, lastName)
// }

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
