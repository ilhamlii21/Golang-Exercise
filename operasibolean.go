package main

import "fmt"

func main() {
	var nilaiAkhit = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhit >= 80
	var lulusaAbsensi bool = absensi >= 79

	var lulus bool = lulusNilaiAkhir && lulusaAbsensi
	fmt.Println("Apakah lulus? ", lulus)

	fmt.Println("Apakah lulus nilai akhir? ", lulusNilaiAkhir)

}
