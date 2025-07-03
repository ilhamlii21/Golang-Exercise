package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("hasil :", hasil)
	} else {
		fmt.Println("error :", err.Error())
	}
}
