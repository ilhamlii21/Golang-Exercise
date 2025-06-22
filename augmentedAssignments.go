package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = 30
	var d = a + b + c
	fmt.Println("Nilai a:", d)

	var i = 10
	i += 10
	i += 10 // i = i + 10
	fmt.Println("Nilai i setelah augmented assignment:", i)

	i += 5 // i = i + 5
	fmt.Println("Nilai i setelah augmented assignment lagi:", i)
}
