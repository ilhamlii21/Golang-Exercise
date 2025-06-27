package main

import "fmt"

//LETSGOOOOOOOOOOOOOOOO RECURSIVEEE

func factorialloop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result

}

func factorialrecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialrecursive(value-1)
	}
}

func main() {
	result := 5 * 4 * 3 * 2
	fmt.Println(result)
	fmt.Println(factorialloop(5))
	fmt.Println(factorialrecursive(4))
}
