package main

import "fmt"

func main() {
	fruit := "carambola"

	if fruit == "carambola" {
		fmt.Println(fruit)
	} else if fruit == "tomate" {
		fmt.Println("tomate")
	} else {
		fmt.Println("Unknown fruit")
	}
}
