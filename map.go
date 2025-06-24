package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Ilham Lii Assidaq"
	// person["address"] = "Jl. Contoh No. 123, Jakarta"

	person := map[string]string{
		"name":    "Ilham Lii Assidaq",
		"address": "Jl. Contoh No. 123, Jakarta",
	}
	fmt.Println(person)            // Menampilkan seluruh elemen dari map
	fmt.Println(person["name"])    // Menampilkan elemen dengan key "name"
	fmt.Println(person["address"]) // Menampilkan elemen dengan key "address"

	book := make(map[string]string) // Membuat map kosong
	book["title"] = "Belajar Go"
	book["author"] = "Ilham Lii Assidaq"
	book["ups"] = "Ups, ada yang salah"

	fmt.Println(book) // Menampilkan seluruh elemen dari map book

	delete(book, "ups") // Menghapus elemen dengan key "ups"

	fmt.Println(book) // Menampilkan seluruh elemen dari map book setelah penghapusan
}
