package main

import "fmt"

func main() {
	name := "Ilh"

	// if name == "Ilham Lii Assidaq" {
	// 	fmt.Println("ini benar")
	// } else {
	// 	fmt.Println("ini salah ")

	// }

	if length := len(name); length > 5 {
		fmt.Println("leng teralu panjang")
	} else {
		fmt.Println("panjangnya pas")
	}

}
