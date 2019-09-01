package main

import "fmt"

func main() {
	birth_year := 1991
	for {
		if birth_year > 2019 {
			break
		}
		fmt.Println(birth_year)
		birth_year++
	}
}
