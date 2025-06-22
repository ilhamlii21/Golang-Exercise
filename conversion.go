package main

import "fmt"

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println("Nilai 32 =", nilai32)
	fmt.Println("Nilai 64 =", nilai64)
	fmt.Println("Nilai 16 =", nilai16)

	var name = "Ilham"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println("Name:", name)
	fmt.Println("First character as uint8:", e)
	fmt.Println("First character as string:", eString)
}
