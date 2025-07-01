package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Cilacap", "Jawa Tengah", "Indonesia"}
	var address2 *Address = &address1 //pointer to address1
	address2.City = "Purbalingga"
	fmt.Println(address1) //ikut berubah
	fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} //address2 now points to a new Address
	fmt.Println(address1)                                      //address1 remains unchanged
	fmt.Println(address2)

}
