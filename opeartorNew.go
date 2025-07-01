package main

import "fmt"

type Address struct {
	City, Province, Country string
} //Struct untuk menyimpan data alamat

func main() {
	var alamat1 *Address = new(Address) //alamat1 adalah pointer ke struct Address
	var alamat2 *Address = alamat1

	alamat2.Country = "Australia" //alamat2 adalah pointer ke struct Address yang sama dengan alamat1
	fmt.Println(alamat1)          //alamat1 tetap sama
	fmt.Println(alamat2)          //alamat2 juga sama

}
