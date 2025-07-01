package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Cilacap", "Jawa Tengah", "Indonesia"}
	// address2 := address1 //copy by value

	// address2.City = "Purwokerto"

	// fmt.Println(address1)
	// fmt.Println(address2)

	address1 := Address{"Cilacap", "Jawa Tengah", "Indonesia"}
	address2 := &address1 //pointer to address1

	address2.City = "Purwokerto"

	fmt.Println(address1) //ikut berubah
	fmt.Println(address2)

}
