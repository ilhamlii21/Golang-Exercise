package main

import "fmt"

func getHello(name string) string {
	hello := "hello " + name
	return hello
}

func main() {
	result := getHello("ilham")
	fmt.Println(result)

	fmt.Println(getHello("Lii"))
	fmt.Println(getHello("Assidaq"))

}
