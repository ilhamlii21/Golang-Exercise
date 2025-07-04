package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked " + name)
	} else {
		fmt.Println("Welcome " + name)

	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Tai"

	}

	registerUser("Ilham", blacklist)
	registerUser("Tai", func(name string) bool {
		return name == "Tai"
	})

}
