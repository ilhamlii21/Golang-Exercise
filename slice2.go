package main

import "fmt"

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) // Menampilkan elemen dari daySlice2 yang sudah ditambahkan "Libur Baru"
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Ilham"
	newSlice[1] = "Ilham"

	fmt.Println(newSlice) // Menampilkan elemen dari newSlice yang berisi "Ilham" dua kalifmt
	fmt.Println(len(newSlice))
	fmt.Println(len(newSlice))

	newSlice2 := append(newSlice, "Ilham Baru")
	fmt.Println(newSlice2) // Menampilkan elemen dari newSlice yang berisi "Ilham" dua kalifmt
	fmt.Println(len(newSlice2))
	fmt.Println(len(newSlice2))

	newSlice2[0] = "Lii"
	fmt.Println(newSlice2) // Menampilkan elemen dari newSlice2 yang sudah diubah
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice) // Menampilkan elemen dari toSlice yang merupakan salinan dari fromSlice
	fmt.Println(toSlice)   // Menampilkan panjang dari toSlice

	iniArray := [...]int{1, 2, 3, 4, 5} // Membuat array dengan {...} syntax
	iniSlice := []int{1, 2, 3, 4}       // Membuat slice dengan [] syntax

	fmt.Println(iniArray) // Menampilkan elemen dari iniArray
	fmt.Println(iniSlice) // Menampilkan elemen dari iniSlice

}
