package main

import "fmt"

func main() {
	var name = "John Doe"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println("Name:", name)
	fmt.Println("First character as uint8:", e)
	fmt.Println("First character as string:", eString)
}
