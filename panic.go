package main

import "fmt"

func endApp() {
	fmt.Println("end app")
	massage := recover()
	fmt.Println("Terjadi Error : ", massage)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}

}

func main() {
	runApp(true)
	fmt.Println("Ilham Lii ")
}
