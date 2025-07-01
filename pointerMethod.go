package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "mr " + man.Name

}
func main() {
	ilham := Man{"Ilham"}
	ilham.Married()
	fmt.Println(ilham.Name)

}
