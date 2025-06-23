package main

import "fmt"

func main() {
	var names [3]string //[3] menunjukan jumlah array yang akan dibuat
	names[0] = "Ilham"  // index dimulai dari 0
	names[1] = "Lii"
	names[2] = "Assidaq"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	var values = [...]int{
		90,
		89,
		78,
		100,
	}

	fmt.Println(values)
	values[0] = 100
	fmt.Println(len(values))
	fmt.Println(len(values))

}
