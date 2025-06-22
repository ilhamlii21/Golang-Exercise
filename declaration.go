package main

import "fmt"

func main() {

	type NoKtp string //deklarasi dengan type data string

	var ktpilham NoKtp = "2311104068"
	var contoh string = "2222222222"

	fmt.Println("No KTP Ilham: ", ktpilham)
	fmt.Println("No KTP Contoh: ", contoh)

}
