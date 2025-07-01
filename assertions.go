package main

import "fmt"

func random() any {
	return true
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt int = result.(int)
	// fmt.Println(resultInt) // This will panic at runtime because result is not an int

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)

	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown type", value)

	}
}
