package main

import (
	"fmt"
	"golang-origin/database"
	_ "golang-origin/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
