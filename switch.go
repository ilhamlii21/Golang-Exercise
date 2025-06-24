package main

import "fmt"

func main() {
	name := "brad"

	// switch name {
	// case "ilham":
	// 	fmt.Println("Hello Ilham")
	// case "brad":
	// 	fmt.Println("Hello Brad")
	// default:
	// 	fmt.Println("Hii")
	// }

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah pas")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length < 5:
		fmt.Println("Nama sudah pas")
	}

}
