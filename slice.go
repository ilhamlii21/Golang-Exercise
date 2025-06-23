package main

import "fmt"

//slice adalah array yang dinamis, bisa berubah-ubah panjangnya
func main() {
	names := [...]string{"Ilham", "Lii", "Assidaq", "Naruto", "Minato", "Itachi"}

	slice1 := names[4:6] // slice dimulai dari index 4 sampai sebelum index 6

	fmt.Println(slice1[0]) // Menampilkan elemen index 0 dari slice
	fmt.Println(slice1[1]) // Menampilkan elemen index 1 dari slice
	fmt.Println(slice1)    // Menampilkan seluruh elemen dari slice

	slice2 := names[:3]
	fmt.Println(slice2) // Menampilkan elemen dari index 0 sampai sebelum index 3

	slice3 := names[3:]
	fmt.Println(slice3) // Menampilkan elemen dari index 3 sampai akhir

	//perbedaan slice dan var adalah slice bisa berubah-ubah panjangnya

	days := [...]string("Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu")

	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) // Menampilkan elemen dari daySlice2 yang sudah ditambahkan "Libur Baru"
	fmt.Println(days)      // Menampilkan elemen dari daySlice1 yang sudah diubah
}
