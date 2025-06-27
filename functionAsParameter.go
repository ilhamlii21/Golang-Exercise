package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Tai" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ilham", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Tai", filter)

}
