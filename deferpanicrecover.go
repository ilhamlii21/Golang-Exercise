package main

import "fmt"

func logging() {
	fmt.Println("Selesai")
}

func runApplication() {
	defer logging() //defer memanggil setelah code dijalankan
	fmt.Println("Run Application...")
}

func main() {
	runApplication()

}
