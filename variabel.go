package main

import "fmt"

// func main() {
// 	var name string

// 	name = "ilham lii"
// 	fmt.Println("Nama saya adalah",  name)

// 	name = "ilham Lii Assidaq"
// 	fmt.Println("Nama saya adalah", name)
// }

//================================================================

// func main() {
// 	var nama = "ilham lii"
// 	fmt.Println("Nama saya adalah", nama)
// }

//================================================================

// func main() {
// 	nam := "ilham lii" // := adalah shorthand untuk deklarasi variabel
// 	fmt.Println("Nama saya adalah", nam)

// 	nam = "ilham Lii Assidaq" // tidak perlu deklarsi ulang dengan var atau :=
// 	fmt.Println("Nama saya adalah", nam)
// }

//================================================================

func main() {
	var (
		FirstName = "Ilham"
		LastName  = "Lii Assidaq"
	)

	fmt.Println("Nama saya adalah", FirstName, LastName)
	fmt.Println("Nama saya adalah", FirstName+" "+LastName)

}
